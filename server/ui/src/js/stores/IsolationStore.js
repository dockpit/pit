var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var IsolationActions = require('../actions/IsolationActions')
var Immutable = require('immutable')
var request = require('superagent')

// private in-memory isolation state
var state = Immutable.Map({
  nrOfIsolations: 0,
  isolations: Immutable.Set(),
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
    case IsolationActions.REFRESH:
      request
        .get('/api/isolations')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          data = JSON.parse(res.text)
          state = state.set('nrOfIsolations', data.length)
          state = state.set('isolations', data)
          IsolationStore.emit(IsolationStore.CHANGED)
        });

      break
    default: 
      return
  }
})

module.exports = IsolationStore