var Dispatcher = require('../dispatcher/AppDispatcher')

var IsolationActions = {
  REFRESH: "REFRESH_ACTION",

  //simply request to reload isolations from server
  refresh: function() {
    Dispatcher.dispatch({
      type: IsolationActions.REFRESH,
    });
  },
}

module.exports = IsolationActions