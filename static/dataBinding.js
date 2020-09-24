import { Observable, Computed, Observale } from "./observable"

export class DataBinding {
    execute(js) {
        return eval(js)
    }

    executeInContext(src, context, attachBindingHelpers = false) {
        if (attachBindingHelpers) {
            context.observable = this.observable
            context.computed = this.computed
            context.bindValue = this.bindValue
        }

        return this.execute.call(context, src)
    }

    observable(value) {
        return new Observale(value)
    }

    computed(calculation, deps) {
        return new Computed(calculation, deps)
    }

    bindValue(input, observable) {
        const initalValue = observable.value
        input.value = initalValue
        observable.subscribe(() => input.value = observable.value)

        let converter = value => value
        if (typeof initalValue === "number") {
            converter = num => isNaN(num = parseFloat(num)) ? 0 : num
        }
        input.onkeyup = () => {
            observable.value = converter(input.value)
        }
    }

    bindAll(elem, context) {
        this.bindLists(elem, context)
        this.bindObservables(elem, context)
    }

    bindObservables(elem, context) {
        const dataBinding = elem.querySelectorAll("[data-bind]");
        dataBinding.forEach(elem => {
            this.bindValue(
                /** @type {HTMLInputElement} */(elem), 
                context[elem.getAttribute("data-bind")]);
        });
    }

    bindLists(elem, context) {
        const listBinding = elem.querySelectorAll("[repeat]");
        listBinding.forEach(elem => {
            const parent = elem.parentElement;
            const expression = elem.getAttribute("repeat");
            elem.removeAttribute("repeat");
            const template = elem.outerHTML;
            parent.removeChild(elem);
            context[expression].forEach(item => {
                let newTemplate = `${template}`;
                const matches = newTemplate.match(/\{\{([^\}]*?)\}\}/g);
                if (matches) {
                    matches.forEach(match => {
                        match = match.replace("{{", "").replace("}}", "");
                        const value = this.executeInContext(`this.${match}`, { item });
                        newTemplate = newTemplate.replace(`{{${match}}}`, value);
                    });
                    parent.innerHTML += newTemplate;
                }
            });
        });
    }
}