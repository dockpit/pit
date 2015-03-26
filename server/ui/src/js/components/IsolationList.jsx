var React = require('react');

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
  	var me = this  	
    return <ul>
		{this.props.isolations.map(function(iso){			
			{if (me.props.selection === iso) {
				return <IsolationItem isSelected={true} key={iso.get('id')} isolation={iso}/> 
			} else{
				return <IsolationItem isSelected={false} key={iso.get('id')} isolation={iso}/> 
			}}
		})}
    </ul>;
  }
});