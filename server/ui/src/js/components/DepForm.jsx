var React = require('react')
var Immutable = require('immutable')

var DepActions  = require('../actions/DepActions')

var TemplateSelector = React.createClass({
	getInitialState: function(){
		return {
			selected: false,
			selectedCat: false
		}
	},

	componentDidMount: function() {
		$(React.findDOMNode(this.refs.accordion)).accordion()
	},

	selectTemplate: function(t, cat) {
		this.setState({
			selected: t.get('id'),
			selectedCat: cat,
		})

		this.props.selectTemplateFn(t)
		React.findDOMNode(this.refs.templateInput).value = t.get('id')
	},

	render: function(){
		var me = this
		var categories = Immutable.Map()

		this.props.templates.forEach(function(t, id){
			var cat = t.get("category")
			if (cat == "") {
				categories = categories.set(id, t)
			} else {
				var existingcat = categories.get(cat)
				if (!existingcat) {
					existingcat = Immutable.Map({name: cat, templates: Immutable.List([])})
				} 				

				var tmpls = existingcat.get('templates')
				existingcat = existingcat.set('templates', tmpls.push(t))

				categories = categories.set(cat, existingcat)
			}
		})

		var items = []
		categories.forEach(function(v, k){
			if (v.get('files')) {
				items.push(<a key={k} onClick={me.selectTemplate.bind(me, v, v.get('id'))} className={'item' + (me.state.selectedCat == v.get('id') ? ' blue active' : '')}>
        				{v.get('name')}
      				</a>)
			} else {
				items.push(<div key={k} className={'item' + (me.state.selectedCat == v.get('name') ? ' blue active' : '')}>
      				<a className="title">
      					<i className="dropdown icon"></i>
        				<em>{v.get('name')}</em>
      				</a>
      				<div className="content">
      					<div className="menu transition hidden">
      						{v.get('templates').map(function(t){
      							return <a 
      							onClick={me.selectTemplate.bind(me, t, v.get('name'))} 
      							key={t.get('id')} 
      							className={'item'+ (me.state.selected == t.get('id') ? ' active' : '')}>
      								{t.get('name')}
      							</a>
      						})}
      					</div>
      				</div>
      			</div>)
			}
		})

		return 	<div ref="accordion" className="ui fluid accordion vertical blue menu">
			{items}

			<input ref="templateInput" name="template" type="hidden" />
		</div>
	}
})

module.exports = React.createClass({
	getInitialState: function() {
		return {
			template: false,
		}
	},

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
			          prompt : 'Please enter a name for the new Dependency'
			        }
			      ]
			    },
				  template: {
			      identifier  : 'template',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please choose a template for the new Dependency'
			        }
			      ]
			    },
			}, {
				onSuccess: me.submit
			})
	},

	selectTemplate: function(t) {
		this.setState({template: t})

		var name = t.get('default_name')
		React.findDOMNode(this.refs.nameInput).value = name
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

		t = this.state.template
		if (!t) {
			return
		}

		var newdep = Immutable.Map({name: name, template: t})
		DepActions.create(newdep)
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
			      Create a new Dependency
			    </div>

			    <div className="content">
					<form ref="form" className="ui form">

						<div className="required field">
							<label>Type</label>
							<TemplateSelector selectTemplateFn={this.selectTemplate} templates={this.props.templates}/>
    					</div>

					    <div className="required field">
					      <label>Name</label>
					      <input ref="nameInput" type="text" placeholder="e.g postgresql, rabbitmq, api.github.com" name="name"/>
					    </div>		    
					</form>
				</div>

				<div className="actions">						
					<button className="ui button" onClick={this.props.closeFormFn}>Cancel</button>		
					<button className="ui button green" onClick={this.submitForm}>Create Dependency</button>
				</div>
			</div>
		</div>
	}
})