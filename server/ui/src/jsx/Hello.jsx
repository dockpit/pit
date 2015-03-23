var React = require('react');

module.exports = React.createClass({
  render: function() {
    return <div>Hello {this.props.name}</div>;
  }
});