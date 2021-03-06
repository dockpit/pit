var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var DepActions = require('../actions/DepActions')
var IsolationActions = require('../actions/IsolationActions')
var Immutable = require('immutable')
var request = require('superagent')

// private in-memory Dep state
var state = Immutable.Map({
  deps: Immutable.List(),
  templates: Immutable.Map(),
})

// maintains and in-memory representation of the Deps
// and accompanying states
var DepStore = assign({}, EventEmitter.prototype, {
  CHANGED: "DEPS_CHANGED",
  DEP_WAS_CREATED: "DEP_WAS_CREATED",

  //return the most up-to-date Dep state
  state: function() {
    return state
  }
})

//register with all actions in the dispatcher
DepStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {
    
    //add a state to a dependency
    case DepActions.ADD_DEP_STATE:
      request
        .post('/api/deps/'+a.args[0].get('id')+'/states')
        .send(a.args[1])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          DepActions.refresh()
          IsolationActions.refresh()
        });

      break;

    //create a new dep on server
    case DepActions.CREATE:
      request
        .post('/api/deps')
        .send(a.args[0])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          DepStore.emit(DepStore.DEP_WAS_CREATED, JSON.parse(res.text))
          
          DepActions.refresh()
        });

      break

    //remove a state from the dep
    case DepActions.REMOVE_DEP_STATE:
      var depid = a.args[0].get('id')
      if(!depid) {
        return console.error("Dep id is empty")
      }

      var sid = a.args[1].get('id')
      if(!sid) {
        return console.error("State id is empty")
      }

      request
        .del('/api/deps/'+depid+'/states/'+sid)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('deps', state.get('deps').map(function(dep){
            if (dep.get('id') == depid) {
              dep = dep.set('states', dep.get('states').filterNot(function(st){
                return st.get('id') == sid
              }))
            }

            return dep
          }))

          IsolationActions.refresh()
        });
      break

    //remove dependency
    case DepActions.REMOVE_DEP:
      var id = a.args[0].get('id')
      if(!id) {
        return console.error("Name is empty")
      }

      request
        .del('/api/deps/'+id)        
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('deps', state.get('deps').filterNot(function(dep){
            return dep.get('id') == id
          }))

          IsolationActions.refresh()
        });
      break

    //refresh depenency state
    case DepActions.REFRESH:
      request
        .get('/api/deps')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          data = JSON.parse(res.text)
          state = state.set('deps', Immutable.fromJS(data))        

          DepStore.emit(DepStore.CHANGED)
        });

      request
        .get('/api/templates')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          data = JSON.parse(res.text)
          state = state.set('templates', Immutable.fromJS(data))
          DepStore.emit(DepStore.CHANGED)
        });

      break
    default: 
      return
  }
})

module.exports = DepStore