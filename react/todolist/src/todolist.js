import React, {Fragment, Component } from 'react';
import Item from './item';

class TodoList extends Component{
  constructor(props){
    super(props);
    this.state = {
      list : [],
      inputValue : ''
    }
    this.handleBtnClick = this.handleBtnClick.bind(this)
    this.handleDeleteItem = this.handleDeleteItem.bind(this)
    this.handleInputChange = this.handleInputChange.bind(this)
  }

  handleBtnClick(){
    if (this.inputValue !== ''){
      console.log("Add an item!")
      this.setState({
        list:[...this.state.list,this.state.inputValue],
        inputValue : ''
      })
    }
  }

  handleInputChange(e){
    this.setState({
      inputValue : e.target.value
    })
  }

  handleDeleteItem(index){
    const list = [...this.state.list]
    list.splice(index,1)
    this.setState({
      list
    })
  }

  getItem(){
    return(
      this.state.list.map(
        (item,index) => {
          return (
            <Item 
            key={index}
            content={item}
            delete={this.handleDeleteItem}
            index={index}
          />
          )
      })
    )
  }
  render(){
    return (
      <Fragment>
        <input value={this.state.inputValue} onChange={this.handleInputChange}/>
        <button onClick={this.handleBtnClick} className="add-btn">add</button>
        <div>{this.getItem()}</div>
      </Fragment>
    )
  }
}

export default TodoList;
