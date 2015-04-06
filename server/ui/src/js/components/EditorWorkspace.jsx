var React = require('react');

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')
var EditorFilePanel  = require('./EditorFilePanel.jsx')
var EditorRunPanel  = require('./EditorRunPanel.jsx')

var EditorCompleteForm = React.createClass({
	preventClickPropagation: function(ev){
		ev.stopPropagation()
	},

	componentDidMount: function() {
		React.findDOMNode(this.refs.nameInput).focus(); 		
	},

	submitForm: function(ev) {
		ev.preventDefault()

		var oldstate = this.props.state
		var newstate = this.props.state.set('name', React.findDOMNode(this.refs.nameInput).value)
		EditorActions.updateState(this.props.depId, oldstate, newstate)

		//optimisticly redirect without waiting
		window.location.pathname = "/"
	},

	render: function(){
		return <div onClick={this.props.closeFormFn} className="ui dimmer modals page transition visible active">
			<div onClick={this.preventClickPropagation} style={{top: '150px'}} className="ui modal transition small visible active">
				<div className="header">
			      Name you newly created state
			    </div>

			    <div className="content">
			    	<p>
			    		It is recommended to give your state a name that
			    		describes the data it contains, {'for'} example:
			    		<ul>
			    			<li><em>two users and a project</em></li>
			    			<li><em>empty database</em></li>
			    		</ul>
			    	</p>
					<form onSubmit={this.submitForm} ref="form" className="ui form">
					    <div className="required field">
					      <label>Name</label>
					      <input 
					      	ref="nameInput" 
					      	type="text" 
					      	defaultValue={this.props.state.get('name')}
					      	name="name"/>
					    </div>		    
					</form>
				</div>

				<div className="actions">						
					<button className="ui button basic" onClick={this.props.closeFormFn}>Cancel</button>		
					<button className="ui button green" onClick={this.submitForm}>Save</button>
				</div>
			</div>
		</div>
	}
})


module.exports = React.createClass({

	getInitialState: function() {
		return {
		  data: EditorStore.state(),
		  showCompleteForm: false,
		}
	},

	componentDidMount: function() {
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)
		EditorActions.refreshState(this.props.depId, this.props.stateId)
	},

	onStoreChange: function() {
		this.setState({data: EditorStore.state()})
	},

	closeCompleteForm: function(ev) {
		ev.preventDefault();
		this.setState({ showCompleteForm: false });
	},

	onCompletedState: function() {
		var runid = this.state.data.get('run').get('id')
		if (runid) {
			EditorActions.stopRun(runid)
		}

		if(this.state.data.get('state').get('name') == EditorStore.DEFAULT_STATE_NAME) {
			this.setState({showCompleteForm: true})
		} else {
			//user changed it already, go to dashboard
			window.location.pathname = "/"
		}
	},

	render: function(){
		return <div className="ui padded grid">
		    <div className="two column row">	
		    	{this.state.showCompleteForm ? <EditorCompleteForm depId={this.props.depId} state={this.state.data.get('state')}  closeFormFn={this.closeCompleteForm}/> : null}

		      	<EditorFilePanel activeFile={this.state.data.get('activeFile')} depId={this.props.depId} state={this.state.data.get('state')} files={this.state.data.get('state').get('files')}/>
				
				<EditorRunPanel 
					depId={this.props.depId} 
					output={this.state.data.get('output')}
					error={this.state.data.get('error')}
					build={this.state.data.get('build')} 
					completedFn={this.onCompletedState}
					run={this.state.data.get('run')} 
					state={this.state.data.get('state')} />		      
		    </div>
		</div>
	}
})