var React = require('react')
var Immutable = require('immutable')

var IsolationActions  = require('../actions/IsolationActions')

module.exports = React.createClass({
	componentDidMount: function() {
		var me = this
		React.findDOMNode(this.refs.nameInput).focus(); 
		$(React.findDOMNode(this.refs.form))
			.form({
				name: {
			      identifier  : 'name',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please enter a name for the new isolation'
			        }
			      ]
			    },
			}, {
				onSuccess: me.submit
			})
	},

	submitForm: function() {
		$(React.findDOMNode(this.refs.form)).form('submit')
	},

	submit: function(ev) {		
		ev.preventDefault();		
		name = React.findDOMNode(this.refs.nameInput).value
		if (name == "") {
			return
		}

		var newiso = Immutable.Map({name: name})
		IsolationActions.create(newiso)
		React.findDOMNode(this.refs.nameInput).value = ''
		this.props.closeFormFn(ev)
	},

	preventClickPropagation: function(ev){
		ev.stopPropagation()
	},

	render: function(){
		return <div onClick={this.props.closeFormFn} className="ui dimmer modals page transition visible active">
			<div onClick={this.preventClickPropagation} style={{top: '150px'}} className="ui modal transition small visible active">
				<div className="header">
			      Create a new Isolation
			    </div>

			    <div className="content">
					<form ref="form" className="ui form">
					    <div className="required field">
					      <label>Name</label>
					      <input ref="nameInput" type="text" placeholder="e.g two users and a project" name="name"/>
					    </div>		    
					</form>
				</div>

				<div className="actions">						
					<button className="ui button" onClick={this.props.closeFormFn}>Cancel</button>		
					<button className="ui button green" onClick={this.submitForm}>Create isolation</button>
				</div>
			</div>
		</div>
	}
})