var assign = require('object-assign')
var EventEmitter = require('events').EventEmitter
var Dispatcher = require('../dispatcher/AppDispatcher')
var IsolationActions = require('../actions/IsolationActions')
var Immutable = require('immutable')
var request = require('superagent')

// private in-memory isolation state
var state = Immutable.Map({
  isolations: Immutable.List(),
  selection: null,
})

// maintains and in-memory representation of the isolations
// and accompanying states
var IsolationStore = assign({}, EventEmitter.prototype, {
  CHANGED: "ISOLATIONS_CHANGED",

  //return the most up-to-date isolation state
  state: function() {
    return state
  },

  sendUpdateReq: function(newiso, oldiso, cb) {
      request
        .put('/api/isolations/'+oldiso.get('id'))
        .send(newiso.toJSON())
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('isolations', state.get('isolations').map(function(curriso){
            if (curriso.get('id') == oldiso.get('id')) {
              return newiso
            }

            return curriso
          }))

          cb()        
        });
  }
})

//register with all actions in the dispatcher
IsolationStore.dispatchToken = Dispatcher.register(function(a){
  switch (a.type) {

    //create a new isolation
    case IsolationActions.CREATE:
      request
        .post('/api/isolations')
        .send(a.args[0])
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          IsolationActions.refresh()
          IsolationStore.emit(IsolationStore.CHANGED)  
        });
      break

    //add a state to an isolation
    case IsolationActions.ADD_STATE:
      var oldiso = a.args[0]
      newiso = oldiso.set('states', oldiso.get('states').set(a.args[1], a.args[2]))
      IsolationStore.sendUpdateReq(newiso, oldiso, function(){
        IsolationStore.emit(IsolationStore.CHANGED)  
      })

      break

    //remove a state from an isolation
    case IsolationActions.REMOVE_STATE:
      var oldiso = a.args[0]
      var id = oldiso.get('id')
      newiso = oldiso.set('states', oldiso.get('states').filterNot(function(v, k){
        return (k == a.args[1].dname && v == a.args[1].sname)       
      }))

      IsolationStore.sendUpdateReq(newiso, oldiso, function(){
        IsolationStore.emit(IsolationStore.CHANGED)  
      })

      break
    
    //isolation name was changed
    case IsolationActions.UPDATE_NAME:
      var id = a.args[1].get('id')
      IsolationStore.sendUpdateReq(a.args[0], a.args[1], function(){
        IsolationStore.emit(IsolationStore.CHANGED)  
      })

      break

    //removes an isolation
    case IsolationActions.REMOVE:
      var id = a.args[0].get('id')
      request
        .del('/api/isolations/'+id)
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('isolations', state.get('isolations').filterNot(function(iso){
            return iso.get('id') == id
          }))

          IsolationStore.emit(IsolationStore.CHANGED)
        });
        
      break
    
    //select a specific isolation
    case IsolationActions.SELECT:
      state = state.set('selection', a.args[0])
      IsolationStore.emit(IsolationStore.CHANGED)
      break

    //refresh list of isolations
    case IsolationActions.REFRESH:
      request
        .get('/api/isolations')
        .end(function(err, res){
          if(err) {
            return console.error(err)
          }

          state = state.set('isolations', Immutable.fromJS(JSON.parse(res.text)))
          IsolationStore.emit(IsolationStore.CHANGED)
        });

      break
    default: 
      return
  }
})

module.exports = IsolationStore