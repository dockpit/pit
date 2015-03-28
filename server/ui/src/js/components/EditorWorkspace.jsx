var React = require('react');

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')
var FilePanel  = require('./EditorFilePanel.jsx')

module.exports = React.createClass({

	getInitialState: function() {
		return {
		  data: EditorStore.state(),
		}
	},

	componentDidMount: function() {
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)
		EditorActions.refreshState(this.props.depName, this.props.stateName)
	},

	onStoreChange: function() {
		this.setState({data: EditorStore.state()})
	},

	render: function(){
		return <div className="ui padded grid">
		    <div className="two column row">

		      <FilePanel activeFile={this.state.data.get('activeFile')} depName={this.props.depName} state={this.state.data.get('state')} files={this.state.data.get('state').get('files')}/>
		      <div className="column">

				<div className="ui two top attached buttons">
		          <button className="ui  button">
		            <i className="repeat icon"></i>(Re)build
		          </button>
		          <button className="ui  button">
		            <i className="terminal icon"></i>Test
		          </button>
		        </div>
				<div className="ui attached segment">			
					<pre>
						Output
					</pre>
				</div>		      	
			    <div className="ui fluid bottom positive attached buttons">
			      <button className="ui button">
			        <i className="checkmark icon"></i>
			        Save
			      </button>
			    </div>
		      </div>
		    </div>
		</div>
	}
})