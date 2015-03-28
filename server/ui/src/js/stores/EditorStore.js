var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var Immutable = require('immutable')
var request = require('superagent')
var EditorActions = require('../actions/EditorActions')

// private in-memory isolation state
var state = Immutable.Map({
  state: Immutable.Map({files: Immutable.Map()}),
  activeFile: '',
})

// maintains various in-memory model representations during editing
var EditorStore = assign({}, EventEmitter.prototype, {
  STATE_CHANGED: "EDITOR_STATE_CHANGED",
  
  //return the most up-to-date isolation state
  state: function() {
    return state
  },

})

EditorStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {

    case EditorActions.UPDATE_FILE_IN_STATE:
      var newfiles = state.get('state').get('files').set(a.args[2], a.args[3])
      var newstate = state.get('state').set('files', newfiles)
      
      request
        .put('/api/deps/'+a.args[0]+'/states/'+a.args[1].get('name'))
        .send(newstate)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('state', newstate)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break;

    //switch to other file
    case EditorActions.REMOVE_FILE_FROM_STATE:
      var newfiles = state.get('state').get('files').delete(a.args[2])
      var newstate = state.get('state').set('files', newfiles)

      request
        .put('/api/deps/'+a.args[0]+'/states/'+a.args[1].get('name'))
        .send(newstate)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('state', newstate)
          state = state.set('activeFile', state.get('state').get('files').keys().next().value)
          
          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break

    //switch to other file
    case EditorActions.SWITCH_FILE:
      state = state.set('activeFile', a.args[0])
      
      EditorStore.emit(EditorStore.STATE_CHANGED)
      break
    
    //new file was added to this state
    case EditorActions.ADD_FILE_TO_STATE:
      var newfiles = state.get('state').get('files').set(a.args[2], '')
      var newstate = state.get('state').set('files', newfiles)
      
      request
        .put('/api/deps/'+a.args[0]+'/states/'+a.args[1].get('name'))
        .send(newstate)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('state', newstate)
          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break;

    //create a new isolation
    case EditorActions.REFRESH_STATE:
      request
        .get('/api/deps/'+a.args[0]+'/states/'+a.args[1])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('state', Immutable.fromJS(JSON.parse(res.text)))
          if (!state.get('activeFile') && state.get('state').get('files').size > 0) {
            state = state.set('activeFile', state.get('state').get('files').keys().next().value) 
          }

          EditorStore.emit(EditorStore.STATE_CHANGED)  
        });

      break

    default: 
      return
  }
})

module.exports = EditorStore