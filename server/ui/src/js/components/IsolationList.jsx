var React = require('react');

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
    return <ul>
		{this.props.isolations.map(function(iso){
			return <IsolationItem key={iso.name} isolation={iso}/> 
		})}
    </ul>;
  }
});