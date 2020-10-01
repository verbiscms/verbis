import { Mark } from 'tiptap'
import { toggleMark } from 'tiptap-commands'

export default class VerbisColour extends Mark {

    get name () {
        return 'textcolor'
    }

    get defaultOptions () {
        return {
            color: ['red'],
        }
    }

    get schema () {
        return {
            attrs: {
                color: {
                    default: 'rgba(0,0,0,1)',
                },
            },
            parseDOM: this.options.color.map(color => ({
                tag: `span[style="color:${color}"]`,
                attrs: { color },
            })),
            toDOM:
                node => {
                    return ['span', {
                        style: `color:${node.attrs.color}`
                    }, 0]
                }
        }
    }

    commands ({ type }) {
        return (attrs) => toggleMark(type, attrs)
    }
}
