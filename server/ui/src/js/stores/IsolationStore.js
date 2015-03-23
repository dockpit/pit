var assign = require('object-assign');
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var IsolationActions = require('../actions/IsolationActions')

// maintains and in-memory representation of the isolations
// and accompanying states
var IsolationStore = assign({}, EventEmitter.prototype, {
  CHANGED: "ISOLATIONS_CHANGED",

  //return the most up-to-date isolation state
  state: function() {

    //@todo return in memory state
    return {}
  }
})

//register with all actions in the dispatcher
Dispatcher.register(function(a){
  switch (a.type) {
    case IsolationActions.REFRESH:
   
      //@todo fetch isolations and update state 
      console.log("Action", a)

      IsolationStore.emit(IsolationStore.CHANGED)
      break
    default: 
      return
  }
})

module.exports = IsolationStore