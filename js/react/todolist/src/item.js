import React,{Component} from 'react';

class Item extends Component {
    constructor(props){
        super(props);
        this.handleDelete = this.handleDelete.bind(this)
    }

    handleDelete(){
        this.props.delete(this.props.index)
    }

    render (){
        const {content} = this.props
        return (<li onClick={this.handleDelete}>{ content }</li>)
    }
}

export default Item;