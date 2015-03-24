var React = require('react');

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
  	var me = this  	
    return <ul>
		{this.props.isolations.map(function(iso){			
			{if (me.props.selection === iso) {
				return <IsolationItem isSelected={true} key={iso.get('name')} isolation={iso}/> 
			} else{
				return <IsolationItem isSelected={false} key={iso.get('name')} isolation={iso}/> 
			}}
		})}
    </ul>;
  }
});