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
})

// maintains and in-memory representation of the Deps
// and accompanying states
var DepStore = assign({}, EventEmitter.prototype, {
  CHANGED: "DEPS_CHANGED",

  //return the most up-to-date Dep state
  state: function() {
    return state
  }
})

//register with all actions in the dispatcher
DepStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {
    case DepActions.REMOVE_DEP_STATE:
      var dname = a.args[0].get('name')
      if(!dname) {
        return console.error("Dep name is empty")
      }

      var sname = a.args[1].get('name')
      if(!sname) {
        return console.error("State name is empty")
      }

      request
        .del('/api/deps/'+dname+'/states/'+sname)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('deps', state.get('deps').map(function(dep){
            if (dep.get('name') == dname) {
              dep = dep.set('states', dep.get('states').filterNot(function(st){
                return st.get('name') == sname
              }))
            }

            return dep
          }))

          IsolationActions.refresh()
          DepStore.emit(DepStore.CHANGED)
        });
      break

    case DepActions.REMOVE_DEP:
      var name = a.args[0].get('name')
      if(!name) {
        return console.error("Name is empty")
      }

      request
        .del('/api/deps/'+name)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('deps', state.get('deps').filterNot(function(dep){
            return dep.get('name') == name
          }))

          IsolationActions.refresh()
          DepStore.emit(DepStore.CHANGED)
        });
      break

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

      break
    default: 
      return
  }
})

module.exports = DepStore