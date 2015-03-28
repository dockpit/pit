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

				{files.map(function(f, i){
					return <EditorFileTab state={me.props.state} depName={me.props.depName} key={i}  activeFile={me.props.activeFile} file={f} switchFileFn={me.switchFile.bind(me, f)}/>
				})}			  
			</div>
			{files.map(function(f, i){				
				return <div key={f.get('name')} className={'ui bottom attached tab segment'+ (f.get('name') == me.props.activeFile ? ' active' : '')}>
					<EditorACE state={me.props.state} depName={me.props.depName} active={f.get('name') == me.props.activeFile ? true : false} file={f} nr={i}/>
				</div>
			})}
		</div>
	}
})