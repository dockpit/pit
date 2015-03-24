var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var DepActions = require('../actions/DepActions')
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