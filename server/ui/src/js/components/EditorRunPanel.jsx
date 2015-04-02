var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

var EditorRunPanel = React.createClass({

	startBuild: function() {
		EditorActions.startBuild(this.props.depId, this.props.state)
	},

	startRun: function() {
		EditorActions.startRun(this.props.depId, this.props.state)
	},

	render: function(){
		var output = this.props.output		
		if(output) {
			output = window.atob(output)
		}		

		var runRunning = (this.props.run.get('id') && !this.props.run.get('is_ready') && !this.props.run.get('error'))

		return <div className="column">

			<div className="ui two steps attached top">
			  <a onClick={this.startBuild} className={'step' + (this.props.build.get('is_running') ? ' disabled': '')}>
			    {this.props.build.get('is_running') ? <div style={{fontSize: '1em'}} className="ui active icon inline loader"></div> : <i className="repeat icon"></i>} 

			    <div className="content">
			      <div className="title">Build</div>
			      <div className="description">Create a Docker Image</div>
			    </div>
			  </a>
			  <a onClick={this.startRun} className={'step' + (!this.props.build.get('image_name') || runRunning  ? ' disabled' : '')}>
			  	{runRunning ? <div style={{fontSize: '1em'}} className="ui active icon inline loader"></div> : <i className="terminal icon"></i>} 
			
			    <div className="content">
			      <div className="title">Test</div>
			      <div className="description">Run the Image</div>
			    </div>
			  </a>
			</div>

			<div className="ui attached bottom segment">			

				{this.props.error ? <div className="ui negative message"><p>{this.props.error}</p></div> : null }

				<pre>
					{output}
				</pre>
			</div>		      	
		</div>
}})

module.exports = EditorRunPanel