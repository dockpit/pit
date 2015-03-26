var React = require('react');

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
  	var me = this  	
    return <div className="ui cards">
		{this.props.isolations.map(function(iso){			
			{if (me.props.selection === iso) {
				return <IsolationItem deps={me.props.deps} isSelected={true} key={iso.get('id')} isolation={iso}/> 
			} else{
				return <IsolationItem deps={me.props.deps} isSelected={false} key={iso.get('id')} isolation={iso}/> 
			}}
		})}
    </div>;
  }
});