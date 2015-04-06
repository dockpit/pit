var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

//throttle function calls
function debounce(fn, delay) {
  var timer = null;
  return function () {
    var context = this, args = arguments;
    clearTimeout(timer);
    timer = setTimeout(function () {
      fn.apply(context, args);
    }, delay);
  };
}

var EditorACE = React.createClass({
	editor: false,
	modeList: false,
	mode: '',

	getInitialState: function(){
		return {
			height: 500,
			lastSavedContent: this.props.file.get('file').get('content')
		}
	},
	
	componentDidMount: function() {
		var me = this
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)

		this.modelist = ace.require("ace/ext/modelist");	    
	    this.mode = this.modelist.getModeForPath(this.props.file.get('name')).mode

	    this.editor = ace.edit(React.findDOMNode(this.refs.editor).id);
	    this.editor.setTheme("ace/theme/github");
    	this.editor.getSession().setMode(this.mode)
    	this.editor.setReadOnly(this.props.file.get('file').get('is_locked'))

		this.editor.getSession().on('change', me.onDirtyChanged);		

		//initial focus
		if(me.props.active) {
			this.editor.focus()						
			this.updateDimensions()
		}

		//add save command
		this.editor.commands.addCommand({
				name: 'saveFile',
				bindKey: {
				win: 'Ctrl-S',
				mac: 'Command-S',
				sender: 'editor|cli'
			},
			exec: function(env, args, request) {
				me.saveFile()
			}
		});

		this.editor.setValue(this.props.file.get('file').get('content'))
		window.addEventListener("resize", this.updateDimensions);
	},

	onDirtyChanged: function() {
			if(this.editor.getValue() == this.state.lastSavedContent) {
				this.props.setDirtyFn(this.props.file, false)
			} else {
				this.props.setDirtyFn(this.props.file, true)
			}
	},

	onStoreChange: function() {		
		this.setState({
			lastSavedContent: this.props.file.get('file').get('content')
		})

		this.onDirtyChanged()
		if (EditorStore.state().get('activeFile') == this.props.file.get('name')) {
			this.editor.focus()
			this.updateDimensions()
		}
	},

    updateDimensions: function() {    	
        this.setState({        
        	height: $(window).height() - 190
        });
    },

    componentWillUnmount: function() {
        window.removeEventListener("resize", this.updateDimensions);
    },

    saveFile: function() {
	    EditorActions.saveFile(
	    	this.props.depId, 
	    	this.props.state, 
	    	this.props.file.get('name'),
	    	this.editor.getValue()
	    )
    },

	render: function(){
		return <div>
			<div style={{float: 'right', zIndex: 100}} className="ui top right attached label">{this.mode.split('/')[2]}</div>
			<div style={{float: 'left', width: '100%', marginTop: '0px !important', height: this.state.height + 'px'}} id={'editor-'+this.props.nr} ref="editor"></div>
			
			{!this.props.file.get('file').get('is_locked') ? <button style={{marginTop: '10px'}} onClick={this.saveFile} className="ui button basic">Save</button> : null}
		</div>
	}
})

var EditorFileTab = React.createClass({
	getInitialState: function() {
		return {
			hover: false
		}
	},

	removeFile: function(ev) {
		if(confirm("Are you sure you want to remove file '"+this.props.file.get('name')+"'?")) {
			EditorActions.removeFileFromState(this.props.depId, this.props.state, this.props.file.get('name'))
		}
		ev.stopPropagation()
	},

	enter: function() { this.setState({hover: true})},
	leave: function() { this.setState({hover: false})},

	render: function() {

		if(this.props.file.get('file').get('is_locked')) {
			return <a style={{opacity: 0.5}} onClick={this.props.switchFileFn}  className={'item'+ (this.props.file.get('name') == this.props.activeFile ? ' active' : '')}>
				{this.props.file.get('name')}
				<i style={{position: 'absolute'}} className="icon lock"></i>
			</a>
		}

		return <a onMouseEnter={this.enter} onMouseLeave={this.leave} onClick={this.props.switchFileFn} className={'item'+ (this.props.file.get('name') == this.props.activeFile ? ' active' : '')}>
			{this.props.file.get('name')}
			{this.props.isDirty ? '*' : null}			
			{this.state.hover ? <i style={{position: 'absolute'}} onClick={this.removeFile} className="icon trash"></i> : null }
		</a>
	}
})

var EditorSettingsTab = React.createClass({
	render: function() {
		return <a onClick={this.props.switchFileFn} className={'item'+ ('__settings' == this.props.activeFile ? ' active' : '')}>
			<i className="icon settings"></i>			
		</a>
	}
})

var EditorPortTable = React.createClass({
	componentDidMount: function() {
		var me = this
		$(React.findDOMNode(this.refs.form))
			.form({
				name: {
			      identifier  : 'cport',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please enter the container port to expose'
			        }
			      ]
			    },
			}, {
				onSuccess: me.newPortBinding
			})
	},

	newPortBinding: function(ev) {
		ev.preventDefault();	

		var cport = React.findDOMNode(this.refs.newContainerPortInput).value
		var hport = React.findDOMNode(this.refs.newHostPortInput).value

		//add the binding
		newbindings = this.props.portBindings.set(cport, Immutable.List([
			Immutable.Map({
				HostIp: "0.0.0.0",
				HostPort: hport
			})
		]))

		this.props.bindingsUpdatedFn(newbindings)

		React.findDOMNode(this.refs.newContainerPortInput).value = ''
		React.findDOMNode(this.refs.newHostPortInput).value = ''
	},

	removePortBinding: function(b, ev) {
		ev.preventDefault()		
		newbindings = this.props.portBindings.remove(b.cport)
		this.props.bindingsUpdatedFn(newbindings)
	},

	onFormSubmit: function(ev) {
		ev.preventDefault()
		$(React.findDOMNode(this.refs.form)).form('submit')
	},

	render: function() {
		var me = this
		var bindings = []
		if (this.props.portBindings) {			

			this.props.portBindings.forEach(function(bs, cport) {
				var b = {
					cport: cport,
				}
				if(bs.size > 0) {
					b["hip"] = bs.get(0).get("HostIp")
					b["hport"] = bs.get(0).get("HostPort")
				}
				
				bindings.push(b)
			})
		}

		return <form ref="form" className="ui form" onSubmit={this.onFormSubmit}>
			<div className="ui error message"></div>
			<table className="ui very basic table">
			  <tbody>
			  	{bindings.map(function(b, i){
			  		return <tr key={i}>
				      <td>{b.cport}</td>
				      <td>{b.hip}:{b.hport}</td>
				      <td>
				      	<a className="ui icon button" onClick={me.removePortBinding.bind(me, b)}>
				      		<i className="minus icon"></i>
				      	</a>
				      </td>
				    </tr>
			  	})}

			    <tr>
			      <td className="field">
			      	<input name="cport" ref="newContainerPortInput" placeholder="Container (e.g 80/tcp)" />
			      </td>
			      <td className="field">
			      	<input ref="newHostPortInput" placeholder="Host (e.g 8000)" />
			      </td>
			      <td>
			      <button className="ui icon button" onClick={this.onFormSubmit}>
			      	<i className="plus icon"></i></button>
			      </td>
			    </tr>

			  </tbody>
			</table>
		</form>
	}
})

var EditorSettings = React.createClass({
	componentDidMount: function() {
		var me = this
		$(React.findDOMNode(this.refs.form))
			.form({
				name: {
			      identifier  : 'name',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please enter a name for the state'
			        }
			      ]
			    },
				ready_pattern: {
			      identifier  : 'ready_pattern',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please choose specify a ready pattern for the state'
			        }
			      ]
			    },
			   ready_timeout: {
			      identifier  : 'ready_timeout',
			      rules: [
			        {
			          type   : 'empty',
			          prompt : 'Please choose specify a ready timeout for the state'
			        }
			      ]
			    },
			}, {
				onSuccess: me.submit
			})
	},

	submit: function(ev) {		
		ev.preventDefault();		

		//get from inputs
		var oldstate = this.props.state
		var newstate = this.props.state.set('name', React.findDOMNode(this.refs.nameInput).value)
		var newsettings = this.props.state.get('settings')
		newsettings = newsettings.set('ready_pattern', React.findDOMNode(this.refs.readyPatternInput).value)
		newsettings = newsettings.set('ready_timeout', React.findDOMNode(this.refs.readyTimeoutInput).value)
		newstate = newstate.set('settings', newsettings)

		EditorActions.updateState(this.props.depId, oldstate, newstate)
	},

	bindingsUpdated: function(newbindings) {
		var newsettings = this.props.state.get('settings')
		var newhostconf = newsettings.get('host_config').set('PortBindings', newbindings)

		EditorActions.changeSettings(newsettings.set('host_config', newhostconf))
	},
	
	validateForm: function() {
		$(React.findDOMNode(this.refs.form)).form('submit')
	},

	render: function() {
		return <div className={'ui bottom attached tab segment '+ ('__settings' == this.props.activeFile ? ' active' : '')}>
			<form ref="form" className="ui form">
				<div className="ui error message"></div>

				  <div className="field">
				    <label>Name</label>
				    <input ref="nameInput" name="name" onChange={this.nameChanged} defaultValue={this.props.state.get('name')} type="text"/>
				  </div>

				  <div className="two fields">
				    <div className="field">
				      <label>Ready Pattern</label>
				      <input ref="readyPatternInput" name="ready_pattern" onChange={this.readyPatternChanged} defaultValue={this.props.state.get('settings').get('ready_pattern')} type="text"/>
				    </div>
				    <div className="field">
				      <label>Ready Timeout</label>
				      <input ref="readyTimeoutInput" name="ready_timeout" onChange={this.readyTimeoutChanged} defaultValue={this.props.state.get('settings').get('ready_timeout')} type="text"/>
				    </div>
				  </div>				  
			</form>
		
			<div className="ui horizontal divider">Ports</div>			  

			<EditorPortTable bindingsUpdatedFn={this.bindingsUpdated} portBindings={this.props.state.get('settings').get('host_config').get('PortBindings')}  />

			<div className="ui divider"></div>
			<div className="ui positive button" onClick={this.validateForm}>
				  <i className="checkmark icon"></i>
				  Save
			</div>	

		</div>
	}
})



module.exports = React.createClass({
	getInitialState: function() {
		return {
			dirtyFiles: {}			
		}
	},

	componentDidMount: function() {
		var me = this
		$(React.findDOMNode(this.refs.addBtn))
		  .popup({
		    inline   : true,
		    on    : 'click',
		    position : 'left center',
		    onVisible: function() {
		    	React.findDOMNode(me.refs.fileNameInput).focus()
		    }
		  })
		;
	},

	switchFile: function(f) {
		EditorActions.switchFile(f.get('name'))
	},

	submitNewFile: function(ev) {
		ev.preventDefault()
		var fname = React.findDOMNode(this.refs.fileNameInput).value
		if (fname) {
			EditorActions.addFileToState(this.props.depId, this.props.state, fname)
		}

		EditorActions.switchFile(fname)
		React.findDOMNode(this.refs.fileNameInput).value = ''
		$(React.findDOMNode(this.refs.addBtn)).popup('hide')
	},

	setDirty: function(f, isDirty) {
		var df = this.state.dirtyFiles
		df[f.get('name')] = isDirty
		this.setState({dirtyFiles: df})
	},

	render: function(){
		var me = this
		var files = Immutable.List()
		this.props.files.forEach(function(f, name){
			files = files.push(Immutable.Map({
				name: name,
				file: f,
			}))
		})

		var settings = Immutable.Map({name: '__settings'})

		return <div className="column"><div className="ui top attached tabular menu">				
				<button ref="addBtn" style={{float: 'right'}} className="ui small basic button"><i className="plus icon"></i>Add file</button>
				<div style={{width: '400px'}} className="ui popup">
					<form onSubmit={this.submitNewFile} className="ui action input">
					  <input style={{border: '1px solid #DFDFDF'}} ref="fileNameInput" type="text" placeholder="e.g script.sh"/>
					  <button type="submit" className="ui icon button">
					    <i className="save icon"></i>
					  </button>
					</form>
				</div>

				<EditorSettingsTab activeFile={me.props.activeFile} switchFileFn={me.switchFile.bind(me, settings)}/>
				{files.map(function(f, i){
					return <EditorFileTab 
						isDirty={me.state.dirtyFiles[f.get('name')]}
						state={me.props.state} 
						depId={me.props.depId} 
						key={i}  
						activeFile={me.props.activeFile} 
						file={f} 
						switchFileFn={me.switchFile.bind(me, f)}/>
				})}			  
			</div>			

			{me.props.activeFile == '__settings' ? <EditorSettings depId={me.props.depId} state={me.props.state} activeFile={me.props.activeFile}/> : null }

			{files.map(function(f, i){				
				return <div key={f.get('name')} className={'ui bottom attached tab segment'+ (f.get('name') == me.props.activeFile ? ' active' : '')}>
					<EditorACE 
						state={me.props.state} 
						setDirtyFn={me.setDirty}
						depId={me.props.depId} 
						active={f.get('name') == me.props.activeFile ? true : false} 
						file={f} 
						nr={i}/>
				</div>
			})}
		</div>
	}
})