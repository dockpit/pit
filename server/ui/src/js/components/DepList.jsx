var React = require('react');

var DepItem = require('./DepItem.jsx')

module.exports = React.createClass({
  render: function() {
    return <ul id="dp-dep-list">
		{this.props.deps.map(function(dep){
			return <DepItem key={dep.get('name')} dep={dep}/> 
		})}
    </ul>;
  }
});