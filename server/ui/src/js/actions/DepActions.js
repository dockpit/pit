var Dispatcher = require('../dispatcher/AppDispatcher')

var DepActions = {
  REFRESH: "DEP_REFRESH_ACTION",
  CREATE: "DEP_CREATE_ACTION",
  REMOVE_DEP: "DEP_REMOVE_ACTION",
  REMOVE_DEP_STATE: "DEP_STATE_REMOVE_ACTION",
  ADD_DEP_STATE: "DEP_STATE_ADD_ACTION",

  //simply request to reload Deps from server
  refresh: function() {
    Dispatcher.dispatch({
      type: DepActions.REFRESH,
    });
  },

  //remove a dependency
  removeDep: function(dep) {
    Dispatcher.dispatch({
      type: DepActions.REMOVE_DEP,
      args: [dep],
    });
  },

  //create a new dependency
  create: function(dep) {
    Dispatcher.dispatch({
      type: DepActions.CREATE,
      args: [dep],
    });
  },

  //remove a dependency's state
  removeDepState: function(dep, state) {
    Dispatcher.dispatch({
      type: DepActions.REMOVE_DEP_STATE,
      args: [dep, state],
    });
  },

  //add a new deps state
  addDepState: function(dep, state) {
    Dispatcher.dispatch({
      type: DepActions.ADD_DEP_STATE,
      args: [dep, state],
    });
  },
}

module.exports = DepActions