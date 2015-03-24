var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var IsolationActions = require('../actions/IsolationActions')
var Immutable = require('immutable')
var request = require('superagent')

// private in-memory isolation state
var state = Immutable.Map({
  isolations: Immutable.List(),
  selection: null,
})

// maintains and in-memory representation of the isolations
// and accompanying states
var IsolationStore = assign({}, EventEmitter.prototype, {
  CHANGED: "ISOLATIONS_CHANGED",

  //return the most up-to-date isolation state
  state: function() {
    return state
  }
})

//register with all actions in the dispatcher
IsolationStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {
    case IsolationActions.REMOVE:
      var name = a.args[0].get('name')
      if(!name) {
        return console.error("Name is empty")
      }

      request
        .del('/api/isolations/'+name)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('isolations', state.get('isolations').filterNot(function(iso){
            return iso.get('name') == name
          }))

          IsolationStore.emit(IsolationStore.CHANGED)
        });
        
      break
    case IsolationActions.SELECT:
      state = state.set('selection', a.args[0])
      IsolationStore.emit(IsolationStore.CHANGED)
      break
    case IsolationActions.REFRESH:
      request
        .get('/api/isolations')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('isolations', Immutable.fromJS(JSON.parse(res.text)))
          IsolationStore.emit(IsolationStore.CHANGED)
        });

      break
    default: 
      return
  }
})

module.exports = IsolationStore