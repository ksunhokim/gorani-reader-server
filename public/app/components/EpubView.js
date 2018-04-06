import React from 'react';
import ReactDOM from 'react-dom';
import Test from '../assets/test.html';

export class EpubView extends React.Component {
  componentDidMount() {
    let ifrm = ReactDOM.findDOMNode(this.refs.iframe);
    ifrm = ifrm.contentWindow || ifrm.contentDocument.document || ifrm.contentDocument;
    ifrm.addEventListener("click",(e)=>{
      console.log(e.target);
    });
  }

  render() {
    return <iframe ref="iframe" src={Test}/>;
  }
};
      