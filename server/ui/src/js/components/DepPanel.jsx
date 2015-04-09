var React = require('react');
var Immutable = require('immutable')

var DepActions  = require('../actions/DepActions')
var StatActions  = require('../actions/StatActions')
var DepStore = require('../stores/DepStore')
var IsolationStore = require('../stores/IsolationStore')
var StatStore = require('../stores/StatStore')
var DepItem = require("./DepItem.jsx")
var DepForm = require('./DepForm.jsx')
var AchievementPopup = require('./AchievementPopup.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: DepStore.state(),
      stats: StatStore.state(),
      isoData: IsolationStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	StatStore.on(StatStore.CHANGED, this.onStoreChange)
    DepStore.on(DepStore.CHANGED, this.onStoreChange)
    DepStore.on(DepStore.DEP_WAS_CREATED, this.onDepWasCreated)
    IsolationStore.on(IsolationStore.CHANGED, this.onStoreChange)    
    
    DepActions.refresh()    
    StatActions.refresh()    
  },

  componentDidUnmount: function() {
    StatStore.removeListener(StatStore.CHANGED, this.onStoreChange)
  	DepStore.removeListener(DepStore.CHANGED, this.onStoreChange)
    DepStore.removeListener(DepStore.DEP_WAS_CREATED, this.onDepWasCreated)
    IsolationStore.removeListener(IsolationStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({
      data: DepStore.state(), 
      isoData: IsolationStore.state(),
      stats: StatStore.state(),
    })

  },

  onDepWasCreated: function(newdep) {
    var depid = newdep.id
    window.location.pathname = "/deps/"+depid+"/states/"+newdep.states[0].id
  },

  openDepForm: function() {
    this.setState({ showForm: true });
  },

  closeDepForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    var me = this
    var me = this   
    var left = Immutable.List()
    var right = Immutable.List()
    this.state.data.get('deps').forEach(function(dep, i){
      if(i & 1) {
          right = right.push(dep)
          return
      } 
      left = left.push(dep)
    })

    var list = <div className="ui two column grid doubling row">
      <div className="column">
        {left.map(function(dep){
          return <DepItem isolations={me.state.isoData.get('isolations')} key={dep.get('id')} dep={dep}/> 
        })}
      </div>
      <div className="column">
        {right.map(function(dep){
          return <DepItem isolations={me.state.isoData.get('isolations')} key={dep.get('id')} dep={dep}/> 
        })}
      </div>
    </div>
    
    //empty list? no grid but a message
    if (this.state.data.get('deps').size < 1) {
      list = <div className="ui attached bottom segment"><div className="ui icon message">
          <i className="database icon"></i>
          <div className="content">
            <div className="header">
              {"You don't have any dependencies configured"}
            </div>
            <p>
              A dependency represents another service you depend on: a databases, HTTP APIs, message queues, etc.
            </p>
            <button className="ui labeled icon button primary" onClick={this.openDepForm}><i className="plus icon">
              </i>Add your first
            </button>
          </div>
        </div></div>
    }

    return <div style={{position: 'relative'}}>
      <h2 className="ui header">
        <img src="/static/src/img/icon_dep@1x.png"/>
        <div className="content">
          Dependencies                   
        </div>
      </h2>

      {this.state.stats.get('stats').get("achievements").size > 0 ? <AchievementPopup achievements={this.state.stats.get('stats').get("achievements")}/> : null}

      {this.state.data.get('deps').size > 0 ? <button style={{position: 'absolute', top: '1.2em', right: '10px'}} className="ui labeled icon button primary" onClick={this.openDepForm}><i className="plus icon"></i>New Dependency</button> : null }

      {this.state.showForm ? <DepForm templates={this.state.data.get('templates')} closeFormFn={this.closeDepForm} dep={this.state.data.get('selection')}/> : null }

      <div className="ui divider"></div>

      {list}
      
    </div>;
  }
});