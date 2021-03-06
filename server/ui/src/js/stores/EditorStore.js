var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var Immutable = require('immutable')
var request = require('superagent')
var EditorActions = require('../actions/EditorActions')

// private in-memory isolation state
var state = Immutable.Map({
  build: Immutable.Map(),
  run: Immutable.Map(),
  state: Immutable.Map({files: Immutable.Map(), settings: Immutable.Map()}),
  output: '',
  error: '',
  activeFile: '',
})

// maintains various in-memory model representations during editing
var EditorStore = assign({}, EventEmitter.prototype, {
  STATE_CHANGED: "EDITOR_STATE_CHANGED",
  DEFAULT_STATE_NAME: "default",
  
  //polls
  interval: null,

  //return the most up-to-date isolation state
  state: function() {
    return state
  },

  sendUpdateReq: function(dname, sid, newstate, cb) {

    request
        .put('/api/deps/'+dname+'/states/'+sid)
        .send(newstate)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          cb()
        });

  }

})

EditorStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {

    //settings changed but should not be send yet
    case EditorActions.EDITOR_CHANGE_SETTINGS:
      var newsettings = a.args[0]
      var newstate = state.get('state').set('settings', newsettings)
      state = state.set('state', newstate)

      EditorStore.emit(EditorStore.STATE_CHANGED) 
      break

    //update
    case EditorActions.UPDATE_STATE:
      var oldstate = a.args[1]
      var newstate = a.args[2]

      EditorStore.sendUpdateReq(a.args[0], oldstate.get('id'), newstate, function() {
          state = state.set('state', newstate)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
      })

      break

    //start build
    case EditorActions.STOP_RUN:
      request
        .del('/api/runs/'+a.args[0])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('run', Immutable.Map())
          EditorStore.emit(EditorStore.STATE_CHANGED)
        })

      break

    //start build
    case EditorActions.START_RUN:

      request
        .post('/api/deps/'+a.args[0]+'/states/'+a.args[1].get('id')+'/runs')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          var run = JSON.parse(res.text)

          if(EditorStore.interval) { clearInterval(EditorStore.interval) }
          EditorStore.interval = setInterval(function(){
              request.get('/api/runs/'+run.id).end(function(err, res){
                if(err) {
                  clearInterval(EditorStore.interval)
                  return console.error(err)
                }                

                var status = JSON.parse(res.text)                
                state = state.set('run', Immutable.fromJS(status))
                state = state.set('output', status.output)              
                state = state.set('error', status.error)                
                EditorStore.emit(EditorStore.STATE_CHANGED)                  

                if(status.error != "") {                  
                  clearInterval(EditorStore.interval)
                  console.error('Run failed:', status.error) 
                }                

                if(status.is_ready === true) {
                  //@todo we might want to continue polling after ready
                  //string is found, this has two complications
                  // - doesn't always seem to work (reaso unknown)
                  // - state.run is reset on which the complete button depends
                  clearInterval(EditorStore.interval)
                }
              })
          }, 1000)
        
          state = state.set('run', Immutable.Map(run))
          EditorStore.emit(EditorStore.STATE_CHANGED)  
        })

      break

    //start build
    case EditorActions.START_BUILD:

      request
        .post('/api/deps/'+a.args[0]+'/states/'+a.args[1].get('id')+'/builds')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          //build was started, continue polling for updates
          var build = JSON.parse(res.text)
          
          if(EditorStore.interval) { clearInterval(EditorStore.interval) }
          EditorStore.interval = setInterval(function(){
              request.get('/api/builds/'+build.id).end(function(err, res){
                if(err) {
                  clearInterval(EditorStore.interval)
                  return console.error(err)
                }

                var status = JSON.parse(res.text)
                state = state.set('build', Immutable.Map(status))
                state = state.set('output', status.output)
                state = state.set('error', status.error)
                EditorStore.emit(EditorStore.STATE_CHANGED)  

                if(status.error != null) {
                  clearInterval(EditorStore.interval)                  
                  console.error('Build failed:', status.error) 
                }                

                //consider done when image name is not empty
                if(status.image_name != '') {
                  clearInterval(EditorStore.interval)  

                  //update state with newly created image
                  newstate = state.get('state').set('image_name', status.image_name)
                  EditorStore.sendUpdateReq(a.args[0], a.args[1].get('id'), newstate, function() {
                      state = state.set('state', newstate)
                      EditorStore.emit(EditorStore.STATE_CHANGED)  
                  })                  
                }                
              })

          }, 100);

          //optimistically assume build will be running soon        
          build.is_running = true
          state = state.set('build', Immutable.Map(build))
          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break;

    //save updated file
    case EditorActions.UPDATE_FILE_IN_STATE:
      var oldfile = state.get('state').get('files').get(a.args[2])
      var newfile = oldfile.set('content', a.args[3])

      var newfiles = state.get('state').get('files').set(a.args[2], newfile)
      var newstate = state.get('state').set('files', newfiles)
      
      //some files where resave, rendering the last build and run out-of-date
      state = state.set('build', Immutable.Map())
      state = state.set('run', Immutable.Map())

      EditorStore.sendUpdateReq(a.args[0], a.args[1].get('id'), newstate, function() {
          state = state.set('state', newstate)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
      })

      break;

    //switch to other file
    case EditorActions.REMOVE_FILE_FROM_STATE:
      var newfiles = state.get('state').get('files').delete(a.args[2])
      var newstate = state.get('state').set('files', newfiles)

      EditorStore.sendUpdateReq(a.args[0], a.args[1].get('id'), newstate, function() {
          state = state.set('state', newstate)
          state = state.set('activeFile', state.get('state').get('files').keys().next().value)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
      })

      break

    //switch to other file
    case EditorActions.SWITCH_FILE:
      state = state.set('activeFile', a.args[0])
      
      EditorStore.emit(EditorStore.STATE_CHANGED)
      break
    
    //new file was added to this state
    case EditorActions.ADD_FILE_TO_STATE:
      var newfiles = state.get('state').get('files').set(a.args[2], Immutable.Map({is_locked: false, content: ''}))
      var newstate = state.get('state').set('files', newfiles)

      EditorStore.sendUpdateReq(a.args[0], a.args[1].get('id'), newstate, function() {
          state = state.set('state', newstate)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
      })
      
      break;

    //create a new isolation
    case EditorActions.REFRESH_STATE:
      request
        .get('/api/deps/'+a.args[0]+'/states/'+a.args[1])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          //set some default values in case server doesnt provide
          var data = JSON.parse(res.text)
          if(!data.settings.host_config.PortBindings) {
              data.settings.host_config.PortBindings = {}
          }

          //set active file to first unlocked file
          state = state.set('state', Immutable.fromJS(data))
          if (!state.get('activeFile') && state.get('state').get('files').size > 0) {
            state.get('state').get('files').forEach(function(f, fname) {
              if (!f.get('is_locked')) {
                state = state.set('activeFile', fname) 
              }
            })
          }

          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break

    default: 
      return
  }
})

module.exports = EditorStore