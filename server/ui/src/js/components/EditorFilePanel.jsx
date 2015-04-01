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
		return {height: 500}
	},
	
	componentDidMount: function() {
		var me = this
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)

		this.modelist = ace.require("ace/ext/modelist");	    
	    this.mode = this.modelist.getModeForPath(this.props.file.get('name')).mode

	    this.editor = ace.edit(React.findDOMNode(this.refs.editor).id);
	    this.editor.setTheme("ace/theme/github");
    	this.editor.getSession().setMode(this.mode)

		this.editor.getSession().on('change', debounce(function(e) {
		    EditorActions.saveFile(
		    	me.props.depName, 
		    	me.props.state, 
		    	me.props.file.get('name'),
		    	me.editor.getValue()
		    )
		},250));		

		//initial focus
		if(me.props.active) {
			this.editor.focus()						
			this.updateDimensions()
		}


		this.editor.setValue(this.props.file.get('content'))
		window.addEventListener("resize", this.updateDimensions);
	},

	onStoreChange: function() {		
		if (EditorStore.state().get('activeFile') == this.props.file.get('name')) {
			this.editor.focus()
			this.updateDimensions()
		}
	},

    updateDimensions: function() {    	
        this.setState({        
        	height: $(window).height() - 140
        });
    },

    componentWillUnmount: function() {
        window.removeEventListener("resize", this.updateDimensions);
    },

	render: function(){
		return <div>
			<div style={{float: 'right', zIndex: 100}} className="ui top right attached label">{this.mode.split('/')[2]}</div>
			<div style={{float: 'left', width: '100%', marginTop: '0px !important', height: this.state.height + 'px'}} id={'editor-'+this.props.nr} ref="editor"></div>
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
			EditorActions.removeFileFromState(this.props.depName, this.props.state, this.props.file.get('name'))
		}
		ev.stopPropagation()
	},

	enter: function() { this.setState({hover: true})},
	leave: function() { this.setState({hover: false})},

	render: function() {
		return <a onMouseEnter={this.enter} onMouseLeave={this.leave} onClick={this.props.switchFileFn} className={'item'+ (this.props.file.get('name') == this.props.activeFile ? ' active' : '')}>
			{this.props.file.get('name')}
			
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

		EditorActions.updateState(this.props.depName, oldstate, newstate)
	},

	render: function() {
		return <form ref="form" className={'ui bottom attached form tab segment'+ ('__settings' == this.props.activeFile ? ' active' : '')}>

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

			  <div className="ui submit button">Save</div>	
		</form>
	}
})



module.exports = React.createClass({
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
			EditorActions.addFileToState(this.props.depName, this.props.state, fname)
		}

		EditorActions.switchFile(fname)
		React.findDOMNode(this.refs.fileNameInput).value = ''
		$(React.findDOMNode(this.refs.addBtn)).popup('hide')
	},

	render: function(){
		var me = this
		var files = Immutable.List()
		this.props.files.forEach(function(content, name){
			files = files.push(Immutable.Map({
				name: name,
				content: content,
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
					return <EditorFileTab state={me.props.state} depName={me.props.depName} key={i}  activeFile={me.props.activeFile} file={f} switchFileFn={me.switchFile.bind(me, f)}/>
				})}			  
			</div>			

			{me.props.activeFile == '__settings' ? <EditorSettings depName={me.props.depName} state={me.props.state} activeFile={me.props.activeFile}/> : null }

			{files.map(function(f, i){				
				return <div key={f.get('name')} className={'ui bottom attached tab segment'+ (f.get('name') == me.props.activeFile ? ' active' : '')}>
					<EditorACE state={me.props.state} depName={me.props.depName} active={f.get('name') == me.props.activeFile ? true : false} file={f} nr={i}/>
				</div>
			})}
		</div>
	}
})