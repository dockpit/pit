var React = require('react');
var Immutable = require('immutable')

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
  	var me = this  	
	return <div style={{boxShadow: "none"}} className="ui styled accordion">
		{this.props.isolations.map(function(iso){
			return <IsolationItem deps={me.props.deps} isSelected={me.props.selection && me.props.selection.get('id') === iso.get('id')} key={iso.get('id')} isolation={iso}/>
		})}
	  </div>

  }
});