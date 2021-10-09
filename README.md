# Instagram Backend API using Go

<h2> APIs </h2>
<h3>Create an User</h3>
<ul>
<li>Should be a POST request</li>
<li>Use JSON request body</li>
<li>URL should be ‘/users'</li>
 </ul>
<h3>Get a user using id</h3>
<ul>
<li>Should be a GET request</li>
<li>Id should be in the url parameter</li>
<li>URL should be ‘/users/<id here>’</li>
  </ul>
<h3>Create a Post</h3>
  <ul>
<li>Should be a POST request</li>
<li>Use JSON request body</li>
<li>URL should be ‘/posts'</li>
  </ul>
<h3>Get a post using id</h3>
  <ul>
<li>Should be a GET request
<li>Id should be in the url parameter</li>
<li>URL should be ‘/posts/<id here>’</li>
  </ul>
<h3>List all posts of a user</h3>
  <ul>
<li>Should be a GET request</li>
<li>URL should be ‘/posts/users/<Id here>'</li>
  </ul>
  
  <h2> Database </h2>
  <h3> MongoDB</h3>
  <p>Users should have the following attributes</p>
    <li>Id </li>
    <li> Name </li>
    <li> Email </li>
    <li> Password </li>

</br>
<p>Posts should have the following Attributes. All fields are mandatory unless marked optional:</p>
    <li> Id </li>
    <li> Caption </li>
    <li> Image URL </li>
    <li> Posted Timestamp </li>
  
