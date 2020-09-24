export class Observale {
    /**
     * 
     * @param {object} value 
     */
    constructor(value) {
        /**
         * Subscriptions
         * @type {ListenerCallback[]}
         */
        this._listeners = []
        /**
         * @type {object}
         */
        this._value = value
    }

    notify() {
        this._listeners.forEach(listener => listener(this._value))
    }

    subscribe(listener) {
        this._listeners.push(listener)
    }

    get value() {
        return this._value
    }

    set value(val) {
        if (val !== this._value) {
            this._value = val
            this.notify()
        }
    }
}

export class Computed extends Observale {
    constructor(value, deps) {
        super(value())
        const listener = () => {
            this._value = value()
            this.notify()
        }
        deps.forEach(dep => dep.subscribe(listener))
    }

    get value() {
        return this._value
    }

    set value(_) {
        throw "Cannot set computed property"
    }
}