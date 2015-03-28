var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

var EditorACE = React.createClass({
	editor: false,

	getInitialState: function(){
		return {height: 500}
	},
	
	componentDidMount: function() {
		var me = this
		EditorStore.on(EditorStore.STATE_CHANGED, this.onStoreChange)
	    
	    this.editor = ace.edit(React.findDOMNode(this.refs.editor).id);
	    this.editor.setTheme("ace/theme/github");
	    this.editor.getSession().setMode("ace/mode/dockerfile");
		this.editor.getSession().on('change', function(e) {
		    //@todo save
		});		

		//initial focus?
		if(me.props.active) {
			this.editor.focus()						
			this.updateDimensions()
		}

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
        	height: $(window).height() - 180
        });
    },

    componentWillUnmount: function() {
        window.removeEventListener("resize", this.updateDimensions);
    },

	render: function(){
		return <div style={{height: this.state.height + 'px'}} id={'editor-'+this.props.nr} ref="editor">
		{this.props.file.get('content')}
		</div>
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

	removeFile: function(f) {
		if(confirm("Are you sure you want to remove file '"+f.get('name')+"'?")) {
			//@todo submit update to state
		}
	},

	submitNewFile: function(ev) {
		ev.preventDefault()
		var fname = React.findDOMNode(this.refs.fileNameInput).value
		if (fname) {
			EditorActions.addFileToState(this.props.depName, this.props.state, fname)
		}

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
					return <a onClick={me.switchFile.bind(me, f)} key={f.get('name')} className={'item'+ (f.get('name') == me.props.activeFile ? ' active' : '')}>
						{f.get('name')}
					</a>
				})}			  
			</div>
			{files.map(function(f, i){				
				return <div key={f.get('name')} className={'ui bottom attached tab segment'+ (f.get('name') == me.props.activeFile ? ' active' : '')}>
					<EditorACE active={f.get('name') == me.props.activeFile ? true : false} file={f} nr={i}/>

				    <div style={{float: 'right'}} className="ui icon buttons">
					  <button onClick={me.removeFile.bind(me, f)} className="ui red button">
					    <i className="trash icon"></i>
					  </button>
					  <button className="ui button">
					    <i className="edit icon"></i>
					  </button>
					</div>
				</div>
			})}
		</div>
	}
})