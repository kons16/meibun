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
            <div>
                {this.state.sentence}<br/>
            </div>
        );
    }
}

export default Books;
