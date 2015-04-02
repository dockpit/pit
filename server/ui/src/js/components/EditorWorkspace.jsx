var React = require('react');

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')
var EditorFilePanel  = require('./EditorFilePanel.jsx')
var EditorRunPanel  = require('./EditorRunPanel.jsx')

module.exports = React.createClass({

	getInitialState: function() {
		return {
		  data: EditorStore.state(),
		}
	},

	componentDidMount: function() {
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)
		EditorActions.refreshState(this.props.depName, this.props.stateId)
	},

	onStoreChange: function() {
		this.setState({data: EditorStore.state()})
	},

	render: function(){
		return <div className="ui padded grid">
		    <div className="two column row">		    	
		      	<EditorFilePanel activeFile={this.state.data.get('activeFile')} depName={this.props.depName} state={this.state.data.get('state')} files={this.state.data.get('state').get('files')}/>
				
				<EditorRunPanel 
					depName={this.props.depName} 
					output={this.state.data.get('output')}
					build={this.state.data.get('build')} 
					run={this.state.data.get('run')} 
					state={this.state.data.get('state')} />		      
		    </div>
		</div>
	}
})