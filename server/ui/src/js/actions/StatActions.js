var Dispatcher = require('../dispatcher/AppDispatcher')

var StatActions = {
  REFRESH: "STAT_REFRESH_ACTION",
  SHOWN: "STAT_SHOWN_ACTION",

  //simply request to reload stats from server
  refresh: function() {
    Dispatcher.dispatch({
      type: StatActions.REFRESH,
    });
  },

  shown: function(achievement) {
    Dispatcher.dispatch({
      type: StatActions.SHOWN,
      args: [achievement],
    });
  }

}

module.exports = StatActions