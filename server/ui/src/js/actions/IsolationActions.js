var Dispatcher = require('../dispatcher/AppDispatcher')

var IsolationActions = {
  REFRESH: "ISOLATION_REFRESH_ACTION",
  SELECT: "ISOLATION_SELECT_ACTION",
  SELECT: "ISOLATION_REMOVE_ACTION",

  //simply request to reload isolations from server
  refresh: function() {
    Dispatcher.dispatch({
      type: IsolationActions.REFRESH,
    });
  },

  //focus a single isolation
  select: function(isolation) {
    Dispatcher.dispatch({
      type: IsolationActions.SELECT,
      args: [isolation],
    });
  },

  //remove a single isolation
  remove: function(isolation) {
    Dispatcher.dispatch({
      type: IsolationActions.REMOVE,
      args: [isolation],
    });
  },
}

module.exports = IsolationActions