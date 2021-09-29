window.render = (function () {
    const TAG = 'p';

    function renderDOM(selector, content) {
        const root = document.querySelector(selector);

        if (!root) {
            return;
        }

        const element = window.dom.createElement(TAG, content);
        root.appendChild(element);
    }

    return {renderDOM};
})();
