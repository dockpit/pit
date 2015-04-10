var React = require('react')
var Immutable = require('immutable')

var DepActions  = require('../actions/DepActions')

var TemplateItem = React.createClass({
	render: function() {
		return <a style={{position: 'relative'}} onClick={this.props.selectTemplateFn} className={'item' + (this.props.selected ? ' blue active' : '')}>
			<img style={{position: 'absolute', left: "0.5em", top: "0.5em"}} height="20" src={tmplIconPath(this.props.template)}/>
			<span style={{paddingLeft: "20px", color: BRAND_BLUE}}>{this.props.template.get('name')}</span>
		</a>
	}
})

var TemplateSelector = React.createClass({
	getInitialState: function(){
		return {
			selected: false,
			selectedCat: false
		}
	},

	componentDidMount: function() {
		$(React.findDOMNode(this.refs.accordion)).accordion()
		$(React.findDOMNode(this.refs.accordion)).accordion("open", 0)
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
				items.unshift(<TemplateItem template={v} key={k} selected={me.state.selectedCat == v.get('id')} selectTemplateFn={me.selectTemplate.bind(me, v, v.get('id'))} />)
			} else {
				items.push(<div key={k} className={'item' + (me.state.selectedCat == v.get('name') ? ' blue active' : '')}>
      				<a className="title">
      					<i style={{left: 0}} className="dropdown icon"></i>
        				<em>{v.get('name')}</em>
      				</a>
      				<div className="content">
      					<div className="menu transition hidden">
      						{v.get('templates').map(function(t){
      							return <TemplateItem 
      							template={t} 
      							key={t.get('id')} 
      							selected={me.state.selected == t.get('id')} 
      							selectTemplateFn={me.selectTemplate.bind(me, t, v.get('name'))} />
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
			<div onClick={this.preventClickPropagation} style={{top: '0px'}} className="ui modal transition small visible active">
				<div style={{paddingLeft: '60px'}} className="header">
				  <img style={{position: 'absolute', top: '5px', left: '5px'}} src="/static/src/img/icon_dep@1x.png"/>
			      Create a new Dependency
			    </div>

			    <div className="content">
					<form ref="form" className="ui form">
						<div className="ui error message"></div>

						<div className="required field">
							<TemplateSelector selectTemplateFn={this.selectTemplate} templates={this.props.templates}/>
    					</div>

					    <div className="required field">
					      <label>Name</label>
					      <input ref="nameInput" type="text" placeholder="e.g postgresql, rabbitmq, api.github.com" name="name"/>
					    </div>		    
					</form>
				</div>

				<div className="actions">						
					<button className="ui basic button" onClick={this.props.closeFormFn}>Cancel</button>		
					<button className="ui button primary" onClick={this.submitForm}>Create Dependency</button>
				</div>
			</div>
		</div>
	}
})