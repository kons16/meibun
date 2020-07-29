import React, { Component } from 'react';

interface BookProps {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
}

interface BookState {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
}

// Book自体のコンポーネント
class Books extends Component<BookProps, BookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            id: props.id,
            sentence: props.sentence,
            title: props.title,
            author: props.author,
            pages: props.pages,
            harts: props.harts
        };
    }

    componentDidMount() {
    }

    render() {
        return (
            <div id="book-component">
                <div id="book-sentence">{this.state.sentence}</div>
                <div id="book-title">『{this.state.title}』</div>
                <div id="book-author">{this.state.author}</div>
                <div id="book-pages">p.{this.state.pages}</div>
                <div id="book-harts">♡{this.state.harts}</div>
            </div>
        );
    }
}

export default Books;
