import { createStore, combineReducers, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'

// Import reducers
import messageReducer from './reducers/message'

const reducer = combineReducers({
  messageReducer
})

const store = createStore(
  reducer,
  applyMiddleware(thunk)
)

export default store