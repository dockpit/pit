var React = require('react');

var IsolationActions  = require('../actions/IsolationActions')
var IsolationStore = require('../stores/IsolationStore')
var DepStore = require('../stores/DepStore')
var IsolationList = require("./IsolationList.jsx")
var IsolationForm = require('./IsolationForm.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: IsolationStore.state(),
      depData: DepStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	IsolationStore.on(IsolationStore.CHANGED, this.onStoreChange)
    DepStore.on(DepStore.CHANGED, this.onStoreChange)
    IsolationActions.refresh()
  },

  componentDidUnmount: function() {
  	IsolationStore.removeListener(IsolationStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({data: IsolationStore.state(), depData: DepStore.state()})
  },

  openIsolationForm: function() {
    this.setState({ showForm: true });
  },

  closeIsolationForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    var disabled = (this.state.depData.get('deps').size < 1)
    if (!disabled && this.state.depData.get('deps').get(0).get('states').size < 1 ) {
      disabled = true
    }

    var list = <IsolationList deps={this.state.depData.get('deps')} selection={this.state.data.get('selection')} isolations={this.state.data.get('isolations')}/>

    return <div className="ui" style={{position: 'relative'}}>
      <h2 style={{paddingLeft: "50px"}} className={'ui top attached header ' + (disabled ? 'disabled': null)}>
        <img style={{position: "absolute", top: "-15px", left: "-10px", "zIndex": 100}} src="/static/src/img/icon_fixture@1x.png" />
        Isolations          
      </h2>

      <button style={{position: 'absolute', top: '10px', right: '10px'}} className={'ui icon button '+ (disabled ? 'disabled': null)} onClick={this.openIsolationForm}><i className="plus icon"></i></button>      
      { this.state.showForm ? <IsolationForm closeFormFn={this.closeIsolationForm} isolation={this.state.data.get('selection')}/> : null }

      <div style={{borderTop: "none", paddingTop: "3px"}} className={'ui bottom attached segment '+ (disabled ? 'disabled': null)}>
        {list}        
      </div>
    </div>;
  }
});