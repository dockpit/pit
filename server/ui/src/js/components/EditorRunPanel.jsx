var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

var EditorRunPanel = React.createClass({

	startBuild: function() {
		EditorActions.startBuild(this.props.depName, this.props.state)
	},

	render: function(){
		return <div className="column">
			<div className="ui two top attached buttons">
			  <button onClick={this.startBuild} className={'ui button' + (this.props.build.get('is_running') ? ' loading disabled': '')}>
				<i className="repeat icon"></i>Build
			  </button>
			  <button className="ui disabled button">
				<i className="play icon"></i>Run
			  </button>
			</div>
			<div className="ui attached bottom segment">			
				<pre>
					{this.props.build.get('output')}
				</pre>
			</div>		      	
		</div>
}})

module.exports = EditorRunPanel