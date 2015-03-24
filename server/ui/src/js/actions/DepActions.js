var Dispatcher = require('../dispatcher/AppDispatcher')

var DepActions = {
  REFRESH: "DEP_REFRESH_ACTION",

  //simply request to reload Deps from server
  refresh: function() {
    Dispatcher.dispatch({
      type: DepActions.REFRESH,
    });
  },
}

module.exports = DepActions