var Dispatcher = require('../dispatcher/AppDispatcher')

var EditorActions = {
  REFRESH_STATE: "EDITOR_REFRESH_ACTION",
  SWITCH_FILE: "EDITOR_SWITCH_FILE",
  ADD_FILE_TO_STATE: "EDITOR_ADD_FILE_TO_STATE_ACTION",

  //simply request to refresh the edited state
  refreshState: function(dname, sname) {
    Dispatcher.dispatch({
      type: EditorActions.REFRESH_STATE,
      args: [dname, sname]
    });
  },

  //adds an file to the given state
  addFileToState: function(dname, state, fname) {
    Dispatcher.dispatch({
      type: EditorActions.ADD_FILE_TO_STATE,
      args: [dname, state, fname]
    });
  },

  //switch to editing another file
  switchFile: function(fname) {
    Dispatcher.dispatch({
      type: EditorActions.SWITCH_FILE,
      args: [fname]
    });
  }

}

module.exports = EditorActions