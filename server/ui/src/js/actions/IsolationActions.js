var Dispatcher = require('../dispatcher/AppDispatcher')

var IsolationActions = {
  REFRESH: "ISOLATION_REFRESH_ACTION",
  CREATE: "ISOLATION_CREATE_ACTION",
  SELECT: "ISOLATION_SELECT_ACTION",
  REMOVE: "ISOLATION_REMOVE_ACTION",
  REMOVE_STATE: "ISOLATION_STATE_REMOVE_ACTION",
  UPDATE_NAME: "ISOLATION_NAME_UPDATE_ACTION",

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
  removeState: function(isolation, state) {
    Dispatcher.dispatch({
      type: IsolationActions.REMOVE_STATE,
      args: [isolation, state],
    });
  },

  //create a new isolation
  create: function(isolation) {
    Dispatcher.dispatch({
      type: IsolationActions.CREATE,
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

  //updateName
  updateName: function(newiso, oldiso) {
    Dispatcher.dispatch({
      type: IsolationActions.UPDATE_NAME,
      args: [newiso, oldiso],
    });
  },
}

module.exports = IsolationActions