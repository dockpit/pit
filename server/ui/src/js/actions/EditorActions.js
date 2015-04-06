var Dispatcher = require('../dispatcher/AppDispatcher')

var EditorActions = {
  CHANGE_SETTINGS: "EDITOR_CHANGE_SETTINGS",
  UPDATE_STATE: "EDITOR_UPDATE_STATE",
  START_BUILD: "EDITOR_START_BUILD",
  START_RUN: "EDITOR_START_RUN",
  STOP_RUN: "EDITOR_STOP_RUN",
  REFRESH_STATE: "EDITOR_REFRESH_ACTION",
  SWITCH_FILE: "EDITOR_SWITCH_FILE",
  UPDATE_FILE_IN_STATE: "EDITOR_UPDATE_FILE_IN_STATE_ACTION",
  ADD_FILE_TO_STATE: "EDITOR_ADD_FILE_TO_STATE_ACTION",
  REMOVE_FILE_FROM_STATE: "EDITOR_REMOVE_FILE_FROM_STATE_ACTION",

  //simply update the whole state
  updateState: function(depid, oldstate, newstate) {
    Dispatcher.dispatch({
      type: EditorActions.UPDATE_STATE,
      args: [depid, oldstate, newstate]
    });
  },

  //updated state settings without sending put
  changeSettings: function(newsettings) {
    Dispatcher.dispatch({
      type: EditorActions.EDITOR_CHANGE_SETTINGS,
      args: [newsettings]
    });
  },

  //simply request to refresh the edited state
  refreshState: function(depid, sid) {
    Dispatcher.dispatch({
      type: EditorActions.REFRESH_STATE,
      args: [depid, sid]
    });
  },

  //remove an file to the given state
  addFileToState: function(depid, state, fname) {
    Dispatcher.dispatch({
      type: EditorActions.ADD_FILE_TO_STATE,
      args: [depid, state, fname]
    });
  },

  //adds an file to the given state
  removeFileFromState: function(depid, state, fname) {
    Dispatcher.dispatch({
      type: EditorActions.REMOVE_FILE_FROM_STATE,
      args: [depid, state, fname]
    });
  },

  //save the current file
  saveFile: function(depid, state, fname, fcontent) {
    Dispatcher.dispatch({
      type: EditorActions.UPDATE_FILE_IN_STATE,
      args: [depid, state, fname, fcontent]
    });
  },

  //Stop any running test containers
  stopRun: function(runId) {
    Dispatcher.dispatch({
      type: EditorActions.STOP_RUN,
      args: [runId]
    });
  },

  //start building the current state
  startBuild: function(depid, state) {
    Dispatcher.dispatch({
      type: EditorActions.START_BUILD,
      args: [depid, state]
    });
  },

  //start running the current state
  startRun: function(depid, state) {
    Dispatcher.dispatch({
      type: EditorActions.START_RUN,
      args: [depid, state]
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