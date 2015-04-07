var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var StatActions = require('../actions/StatActions')

var Immutable = require('immutable')
var request = require('superagent')

var state = Immutable.Map({
  stats: Immutable.fromJS({achievements: {}}),
})

// maintains and in-memory representation of the Deps
// and accompanying states
var StatStore = assign({}, EventEmitter.prototype, {
  CHANGED: "STATS_CHANGED",

  //return the most up-to-date Dep state
  state: function() {
    return state
  }
})

//register with all actions in the dispatcher
StatStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {
    case StatActions.SHOWN:
      
      var newach = state.get('stats').get('achievements').map(function(ach){
        if (a.args[0].get('id') == ach.get('id')) {
          return ach.set('shown', true)
        }

        return ach
      })

      var newstats = state.get('stats').set('achievements', newach)
      request
        .put('/api/stats')
        .send(newstats.toJS())
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('stats', Immutable.fromJS(JSON.parse(res.text)))
          StatStore.emit(StatStore.CHANGED)
        });

      break

    case StatActions.REFRESH:
      request
        .get('/api/stats')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('stats', Immutable.fromJS(JSON.parse(res.text)))
          StatStore.emit(StatStore.CHANGED)
        });

      break
  }
})

module.exports = StatStore