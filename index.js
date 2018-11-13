{

  const gates = new API('https://gates.codehut.me', { credentials: 'omit' });

  class Header extends Jinkela {
    get template() {
      return `
        <header>
          <button type="button" on-click="{run}">Run</button>
          <button type="button" on-click="{runFib}">Fibonacci</button>
        </header>
      `;
    }

    get run() {
      return () => {
        this.element.dispatchEvent(new CustomEvent('run', { bubbles: true }));
      };
    }

    get runFib() {
      const code = `(function () {
  let fib = function (n) {
    return n <= 1 && 1 || fib(n - 2) + fib(n - 1);
  };
  return fib;
})()(10)
`;
      return () => {
        this.element.dispatchEvent(new CustomEvent('run', { bubbles: true, detail: { code } }));
      };
    }

    get styleSheet() {
      return `
        :scope {
          height: 30px;
        }
      `;
    }
  }

  class Editor extends Jinkela {
    get template() {
      return `
        <div></div>
      `;
    }

    init() {
      this.editor = monaco.editor.create(this.element, {
        automaticLayout: true
      });
    }

    get value() {
      return this.editor.getModel().getValue();
    }

    set value(value) {
      this.editor.getModel().setValue(value);
    }

    get styleSheet() {
      return `
        :scope {
          flex: 2;
          min-height: 0;
        }
      `;
    }
  }

  class Result extends Jinkela {
    get template() {
      return `
        <pre></pre>
      `;
    }

    update(result) {
      this.element.textContent = result;
    }

    updateError(message) {
      this.element.textContent = message;
    }

    get styleSheet() {
      return `
        :scope {
          flex: 1;
          min-height: 0;
          padding: 10px;
        }
      `;
    }
  }

  class Main extends Jinkela {
    get Header() { return Header; }
    get Editor() { return Editor; }
    get Result() { return Result; }

    get template() {
      return `
        <div on-run="{run}">
          <jkl-header></jkl-header>
          <jkl-editor ref="editor"></jkl-editor>
          <jkl-result ref="result"></jkl-result>
        </div>
      `;
    }

    get run() {
      return async (e) => {
        if (e.detail) {
          const { code } = e.detail;
          this.editor.value = code;
        }
        const code = this.editor.value;
        try {
          const { result } = await gates.run.post({ body: { code } });
          this.result.update(result);
        } catch (error) {
          this.result.updateError(error.message);
        }
      };
    }

    get styleSheet() {
      return `
        html, body {
          height: 100%;
          margin: 0;
          overflow: hidden;
        }

        :scope {
          display: flex;
          flex-direction: column;
          height: 100%;
        }
      `;
    }
  }

  require.config({ paths: { 'vs': 'https://cdn.jsdelivr.net/npm/monaco-editor@0.15.1/min/vs' }});
  window.MonacoEnvironment = {
    getWorkerUrl: function(workerId, label) {
      return `data:text/javascript;charset=utf-8,${encodeURIComponent(`
        self.MonacoEnvironment = {
          baseUrl: 'https://cdn.jsdelivr.net/npm/monaco-editor@0.15.1/min/'
        };
        importScripts('https://cdn.jsdelivr.net/npm/monaco-editor@0.15.1/min/vs/base/worker/workerMain.js');`
      )}`;
    }
  };

  require(["vs/editor/editor.main"], function () {
    new Main().to(document.body);
  });

}
