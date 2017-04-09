import Vue from 'vue'
import Vuex from 'vuex'
import {Rpc} from './rpc.js'

Vue.use(Vuex)

function createState () {
  return {
    selectedRoot: '',
    selectedFiles: [],
    path: [],
    loading: true
  }
};

const state = {
  selectedState: '0',
  roots: [],
  states: [],
  uiState: 'initializing'
}

const actions = {
  init: function ({commit}) {
    Rpc.getRoots()
      .then((response) => {
        console.log(response)
        commit('init', response.result)
      })
  }
}

const otherStateId = (id) => {
  return id === '0'
    ? '1'
    : '0'
}

const getPathString = path => {
  return path.reduce((acc, p) => {
    return acc + '/' + p
  }, '')
}

const getters = {
  currentState (state) {
    return state.states[state.selectedState]
  },
  currentPathString (state) {
    let path = state.states[state.selectedState].path
    return getPathString(path)
  },
  otherState (state) {
    return state.states[otherStateId(state.selectedState)]
  },
  otherPathString (state) {
    let otherState = state.states[otherStateId(state.selectedState)]
    return getPathString(otherState.path)
  }
}
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    init (state, roots) {
      state.states = [createState(), createState()]
      state.roots = roots
      state.uiState = 'browse'
      state.loading = false
      if (roots.length > 0) {
        state.states[0].selectedRoot = state.roots[0]
        state.states[1].selectedRoot = roots.length > 1
          ? state.roots[1]
          : state.roots[0]
      }
    },
    selectRoot (state, message) {
      var viewState = state.states[message.stateId]
      viewState.selectedRoot = message.value
      viewState.path = []
    },
    selectSingleFile (state, message) {
      if (state.selectedState !== message.stateId) {
        state.states[state.selectedState].selectedFiles = []
      }
      state.states[message.stateId].selectedFiles = [message.value]
    },
    selectFile (state, message) {
      if (state.selectedState !== message.stateId) {
        state.states[state.selectedState].selectedFiles = []
      }

      let selectedFiles = state.states[message.stateId].selectedFiles
      let fileIndex = selectedFiles.indexOf(message.value)
      if (fileIndex === -1) {
        selectedFiles.push(message.value)
      } else {
        selectedFiles.splice(fileIndex, 1)
      }
    },
    selectView (state, viewId) {
      if (state.selectedState !== viewId) {
        state.states[state.selectedState].selectedFiles = []
      }
      state.selectedState = viewId
    },
    changePath (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.push(message.value)
    },
    changePathToParent (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.pop()
    },
    setPath (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.splice(message.value, viewState.path.length)
    },
    rename (state) {
      state.uiState = 'rename'
    },
    renameWait (state) {
      state.uiState = 'rename-wait'
    },
    mkdir (state) {
      state.uiState = 'mkdir'
    },
    mkdirWait (state) {
      state.uiState = 'mkdir-wait'
    },
    copyWait (state) {
      state.uiState = 'copy-wait'
    },
    moveWait (state) {
      state.uiState = 'move-wait'
    },
    deleteFile (state) {
      state.uiState = 'delete-file'
    },
    deleteFileWait (state) {
      state.uiState = 'delete-file-wait'
    },
    browse (state) {
      state.uiState = 'browse'
    },
    error (state, error) {
      state.uiState = 'error'
      state.error = error
    },
    commandFinished (state) {
      state.uiState = 'browse'
      state.states[state.selectedState].selectedFiles = []
    }
  },
  actions: actions
})
