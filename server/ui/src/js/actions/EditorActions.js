var Dispatcher = require('../dispatcher/AppDispatcher')

var EditorActions = {
  UPDATE_STATE: "EDITOR_UPDATE_STATE",
  START_BUILD: "EDITOR_START_BUILD",
  REFRESH_STATE: "EDITOR_REFRESH_ACTION",
  SWITCH_FILE: "EDITOR_SWITCH_FILE",
  UPDATE_FILE_IN_STATE: "EDITOR_UPDATE_FILE_IN_STATE_ACTION",
  ADD_FILE_TO_STATE: "EDITOR_ADD_FILE_TO_STATE_ACTION",
  REMOVE_FILE_FROM_STATE: "EDITOR_REMOVE_FILE_FROM_STATE_ACTION",

  //simply update the whole state
  updateState: function(dname, oldstate, newstate) {
    Dispatcher.dispatch({
      type: EditorActions.UPDATE_STATE,
      args: [dname, oldstate, newstate]
    });
  },

  //simply request to refresh the edited state
  refreshState: function(dname, sid) {
    Dispatcher.dispatch({
      type: EditorActions.REFRESH_STATE,
      args: [dname, sid]
    });
  },

  //remove an file to the given state
  addFileToState: function(dname, state, fname) {
    Dispatcher.dispatch({
      type: EditorActions.ADD_FILE_TO_STATE,
      args: [dname, state, fname]
    });
  },

  //adds an file to the given state
  removeFileFromState: function(dname, state, fname) {
    Dispatcher.dispatch({
      type: EditorActions.REMOVE_FILE_FROM_STATE,
      args: [dname, state, fname]
    });
  },

  //save the current file
  saveFile: function(dname, state, fname, fcontent) {
    Dispatcher.dispatch({
      type: EditorActions.UPDATE_FILE_IN_STATE,
      args: [dname, state, fname, fcontent]
    });
  },

  //start building the current state
  startBuild: function(dname, state) {
    Dispatcher.dispatch({
      type: EditorActions.START_BUILD,
      args: [dname, state]
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