package main

type HTTPCode struct {
	Code        int
	Name        string
	Description string
}

var HTTPCodes []HTTPCode = []HTTPCode{
	{
		Code: 100,
		Name: "Continue",
		Description: `The initial part of a request has been received and has not yet been rejected by the server. The server intends to send a final response after the request has been fully received and acted upon.

When the request contains an Expect header field that includes a 100-continue expectation, the 100 response indicates that the server wishes to receive the request payload body. The client ought to continue sending the request and discard the 100 response.

If the request did not contain an Expect header field containing the 100-continue expectation, the client can simply discard this interim response.`,
	},
	{
		Code: 101,
		Name: "Switching Protocols",
		Description: `

The server understands and is willing to comply with the client&#39;s request, via the Upgrade header field<sup><a href="#ref-1">1</a></sup>, for a change in the application protocol being used on this connection.

The server MUST generate an Upgrade header field in the response that indicates which protocol(s) will be switched to immediately after the empty line that terminates the 101 response.

It is assumed that the server will only agree to switch protocols when it is advantageous to do so. For example, switching to a newer version of HTTP might be advantageous over older versions, and switching to a real-time, synchronous protocol might be advantageous when delivering resources that use such features.
`,
	},
	{
		Code: 102,
		Name: "Processing",
		Description: `

An interim response used to inform the client that the server has accepted the complete request, but has not yet completed it.

This status code SHOULD only be sent when the server has a reasonable expectation that the request will take significant time to complete. As guidance, if a method is taking longer than 20 seconds (a reasonable, but arbitrary value) to process the server SHOULD return a 102 (Processing) response. The server MUST send a final response after the request has been completed.

Methods can potentially take a long period of time to process, especially methods that support the Depth header. In such cases the client may time-out the connection while waiting for a response. To prevent this the server may return a 102 Processing status code to indicate to the client that the server is still processing the method.
`,
	},
	{
		Code: 200,
		Name: "OK",
		Description: `

The request has succeeded.

The payload sent in a 200 response depends on the request method. For the methods defined by this specification, the intended meaning of the payload can be summarized as:
<ul>
<li><code>GET</code> a representation of the target resource</li>
<li><code>HEAD</code> the same representation as <code>GET</code>, but without the representation data</li>
<li><code>POST</code> a representation of the status of, or results obtained from, the action;<ul>
<li><code>PUT</code> <code>DELETE</code> a representation of the status of the action;</li>
<li><code>OPTIONS</code> a representation of the communications options;</li>
<li><code>TRACE</code> a representation of the request message as received by the end server.</li>
</ul>
</li>
</ul>

Aside from responses to CONNECT, a 200 response always has a payload, though an origin server MAY generate a payload body of zero length. If no payload is desired, an origin server ought to send <a href="/204">204 No Content</a> instead.  For CONNECT, no payload is allowed because the successful result is a tunnel, which begins immediately after the 200 response header section.

A 200 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 201,
		Name: "Created",
		Description: `

The request has been fulfilled and has resulted in one or more new resources being created.

The primary resource created by the request is identified by either a Location header field in the response or, if no Location field is received, by the effective request URI.

The 201 response payload typically describes and links to the resource(s) created. See <a href="https://tools.ietf.org/html/rfc7231#section-7.2">Section 7.2 of RFC7231</a> for a discussion of the meaning and purpose of validator header fields, such as ETag and Last-Modified, in a 201 response.
`,
	},
	{
		Code: 202,
		Name: "Accepted",
		Description: `

The request has been accepted for processing, but the processing has not been
completed. The request might or might not eventually be acted upon, as it might
be disallowed when processing actually takes place.

There is no facility in HTTP for re-sending a status code from an asynchronous
operation.

The 202 response is intentionally noncommittal. Its purpose is to allow a server to accept a request for some other process (perhaps a batch-oriented process that is only run once per day) without requiring that the user agent&#39;s connection to the server persist until the process is completed. The representation sent with this response ought to describe the request&#39;s current status and point to (or embed) a status monitor that can provide the user with an estimate of when the request will be fulfilled.
`,
	},
	{
		Code: 203,
		Name: "Non-authoritative Information",
		Description: `

The request was successful but the enclosed payload has been modified from that of the origin server&#39;s <a href="/200">200 OK</a> response by a transforming proxy<sup><a href="#ref-1">1</a></sup>.

This status code allows the proxy to notify recipients when a transformation has been applied, since that knowledge might impact later decisions regarding the content. For example, future cache validation requests for the content might only be applicable along the same request path (through the same proxies).

The 203 response is similar to the Warning code of 214 Transformation Applied<sup><a href="#ref-2">2</a></sup>, which has the advantage of being applicable to responses with any status code.

A 203 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-3">3</a></sup>.
`,
	},
	{
		Code: 204,
		Name: "No Content",
		Description: `

The server has successfully fulfilled the request and that there is no additional content to send in the response payload body.

Metadata in the response header fields refer to the target resource and its selected representation after the requested action was applied.

For example, if a 204 status code is received in response to a PUT request and the response contains an ETag header field, then the PUT was successful and the ETag field-value contains the entity-tag for the new representation of that target resource.

The 204 response allows a server to indicate that the action has been successfully applied to the target resource, while implying that the user agent does not need to traverse away from its current &quot;document view&quot; (if any). The server assumes that the user agent will provide some indication of the success to its user, in accord with its own interface, and apply any new or updated metadata in the response to its active representation.

For example, a 204 status code is commonly used with document editing interfaces corresponding to a &quot;save&quot; action, such that the document being saved remains available to the user for editing. It is also frequently used with interfaces that expect automated data transfers to be prevalent, such as within distributed version control systems.

A 204 response is terminated by the first empty line after the header fields because it cannot contain a message body.

A 204 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 205,
		Name: "Reset Content",
		Description: `

The server has fulfilled the request and desires that the user agent reset the &quot;document view&quot;, which caused the request to be sent, to its original state as received from the origin server.

This response is intended to support a common data entry use case where the user receives content that supports data entry (a form, notepad, canvas, etc.), enters or manipulates data in that space, causes the entered data to be submitted in a request, and then the data entry mechanism is reset for the next entry so that the user can easily initiate another input action.

Since the 205 status code implies that no additional content will be provided, a server MUST NOT generate a payload in a 205 response. In other words, a server MUST do one of the following for a 205 response: a) indicate a zero-length body for the response by including a Content-Length header field with a value of 0; b) indicate a zero-length payload for the response by including a Transfer-Encoding header field with a value of chunked and a message body consisting of a single chunk of zero-length; or, c) close the connection immediately after sending the blank line terminating the header section.
`,
	},
	{
		Code: 206,
		Name: "Partial Content",
		Description: `

The server is successfully fulfilling a range request for the target resource by transferring one or more parts of the selected representation that correspond to the satisfiable ranges found in the request&#39;s Range header field<sup><a href="#ref-1">1</a></sup>.

If a single part is being transferred, the server generating the 206 response MUST generate a Content-Range header field, describing what range of the selected representation is enclosed, and a payload consisting of the range. For example:
<pre><code>HTTP/1.1 206 Partial Content
Date: Wed, 15 Nov 1995 06:25:24 GMT
Last-Modified: Wed, 15 Nov 1995 04:58:08 GMT
Content-Range: bytes 21010-47021/47022
Content-Length: 26012
Content-Type: image/gif

... 26012 bytes of partial image data ...
</code></pre>
If multiple parts are being transferred, the server generating the 206 response MUST generate a &quot;multipart/byteranges&quot; payload<sup><a href="#ref-2">2</a></sup>, and a Content-Type header field containing the multipart/byteranges media type and its required boundary parameter. To avoid confusion with single-part responses, a server MUST NOT generate a Content-Range header field in the HTTP header section of a multiple part response (this field will be sent in each part instead).

Within the header area of each body part in the multipart payload, the server MUST generate a Content-Range header field corresponding to the range being enclosed in that body part. If the selected representation would have had a Content-Type header field in a <a href="/200">200 OK</a> response, the server SHOULD generate that same Content-Type field in the header area of each body part. For example:
<pre><code>HTTP/1.1 206 Partial Content
Date: Wed, 15 Nov 1995 06:25:24 GMT
Last-Modified: Wed, 15 Nov 1995 04:58:08 GMT
Content-Length: 1741
Content-Type: multipart/byteranges; boundary=THIS_STRING_SEPARATES

--THIS_STRING_SEPARATES
Content-Type: application/pdf
Content-Range: bytes 500-999/8000

...the first range...
--THIS_STRING_SEPARATES
Content-Type: application/pdf
Content-Range: bytes 7000-7999/8000

...the second range
--THIS_STRING_SEPARATES--
</code></pre>
When multiple ranges are requested, a server MAY coalesce any of the ranges that overlap, or that are separated by a gap that is smaller than the overhead of sending multiple parts, regardless of the order in which the corresponding byte-range-spec appeared in the received Range header field. Since the typical overhead between parts of a multipart/byteranges payload is around 80 bytes, depending on the selected representation&#39;s media type and the chosen boundary parameter length, it can be less efficient to transfer many small disjoint parts than it is to transfer the entire selected representation.

A server MUST NOT generate a multipart response to a request for a single range, since a client that does not request multiple parts might not support multipart responses. However, a server MAY generate a multipart/byteranges payload with only a single body part if multiple ranges were requested and only one range was found to be satisfiable or only one range remained after coalescing. A client that cannot process a multipart/byteranges response MUST NOT generate a request that asks for multiple ranges.

When a multipart response payload is generated, the server SHOULD send the parts in the same order that the corresponding byte-range-spec appeared in the received Range header field, excluding those ranges that were deemed unsatisfiable or that were coalesced into other ranges. A client that receives a multipart response MUST inspect the Content-Range header field present in each body part in order to determine which range is contained in that body part; a client cannot rely on receiving the same ranges that it requested, nor the same order that it requested.

When a 206 response is generated, the server MUST generate the following header fields, in addition to those required above, if the field would have been sent in a <a href="/200">200 OK</a> response to the same request: Date, Cache-Control, ETag, Expires, Content-Location, and Vary.

If a 206 is generated in response to a request with an If-Range header field, the sender SHOULD NOT generate other representation header fields beyond those required above, because the client is understood to already have a prior response containing those header fields. Otherwise, the sender MUST generate all of the representation header fields that would have been sent in a <a href="/200">200 OK</a> response to the same request.

A 206 response is cacheable by default; i.e., unless otherwise indicated by explicit cache controls<sup><a href="#ref-3">3</a></sup>.
`,
	},
	{
		Code: 207,
		Name: "Multi-Status",
		Description: `

A Multi-Status response conveys information about multiple resources in situations where multiple status codes might be appropriate.

The default Multi-Status response body is a text/xml or application/xml HTTP entity with a &#39;multistatus&#39; root element. Further elements contain 200, 300, 400, and 500 series status codes generated during the method invocation. 100 series status codes SHOULD NOT be recorded in a &#39;response&#39; XML element.

Although &#39;207&#39; is used as the overall response status code, the recipient needs to consult the contents of the multistatus response body for further information about the success or failure of the method execution. The response MAY be used in success, partial success and also in failure situations.

The &#39;multistatus&#39; root element holds zero or more &#39;response&#39; elements in any order, each with information about an individual resource. Each &#39;response&#39; element MUST have an &#39;href&#39; element to identify the resource.

A Multi-Status response uses one out of two distinct formats for representing the status:

1. A &#39;status&#39; element as child of the &#39;response&#39; element indicates the status of the message execution for the identified resource as a whole<sup><a href="#ref-1">1</a></sup>. Some method definitions provide information about specific status codes clients should be prepared to see in a response. However, clients MUST be able to handle other status codes, using the generic rules defined in <a href="https://tools.ietf.org/html/rfc2616#section-10">RFC2616 Section 10</a>.

2. For PROPFIND and PROPPATCH, the format has been extended using the &#39;propstat&#39; element instead of &#39;status&#39;, providing information about individual properties of a resource.  This format is specific to PROPFIND and PROPPATCH, and is described in detail in <a href="https://tools.ietf.org/html/rfc4918#section-9.1">RFC4918 Section 9.1</a> and <a href="https://tools.ietf.org/html/rfc4918#section-9.2">RFC4918 Section 9.2</a>.
`,
	},
	{
		Code: 208,
		Name: "Already Reported",
		Description: `

Used inside a DAV: propstat response element to avoid enumerating the internal members of multiple bindings to the same collection repeatedly.

For each binding to a collection inside the request&#39;s scope, only one will be reported with a 200 status, while subsequent DAV:response elements for all other bindings will use the 208 status, and no DAV:response elements for their descendants are included.

Note that the 208 status will only occur for &quot;Depth: infinity&quot; requests, and that it is of particular importance when the multiple collection bindings cause a bind loop<sup><a href="#ref-1">1</a></sup>.

A client can request the DAV:resource-id property in a PROPFIND request to guarantee that they can accurately reconstruct the binding structure of a collection with multiple bindings to a single resource.

For backward compatibility with clients not aware of the 208 status code appearing in multistatus response bodies, it SHOULD NOT be used unless the client has signaled support for this specification using the &quot;DAV&quot; request header<sup><a href="#ref-2">2</a></sup>. Instead, a <a href="/508">508 Loop Detected</a> status should be returned when a binding loop is discovered. This allows the server to return the 508 as the top-level return status, if it discovers it before it started the response, or in the middle of a multistatus, if it discovers it in the middle of streaming out a multistatus response.
`,
	},
	{
		Code: 226,
		Name: "IM Used",
		Description: `

The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.

The actual current instance might not be available except by combining this response with other previous or future responses, as appropriate for the specific instance-manipulation(s). If so, the headers of the resulting instance are the result of combining the headers from the 226 response and the other instances, following the rules in <a href="https://tools.ietf.org/html/rfc2616#section-13.5.3">section 13.5.3</a> of the HTTP/1.1 specification.

The request MUST have included an A-IM header field listing at least one instance-manipulation. The response MUST include an Etag header field giving the entity tag of the current instance.

A response received with a status code of 226 MAY be stored by a cache and used in reply to a subsequent request, subject to the HTTP expiration mechanism and any Cache-Control headers, and to the requirements in <a href="https://tools.ietf.org/html/rfc3229#section-10.6">section 10.6</a>.

A response received with a status code of 226 MAY be used by a cache, in conjunction with a cache entry for the base instance, to create a cache entry for the current instance.
`,
	},
	{
		Code: 300,
		Name: "Multiple Choices",
		Description: `

The target resource has more than one representation, each with its own more specific identifier, and information about the alternatives is being provided so that the user (or user agent) can select a preferred representation by redirecting its request to one or more of those identifiers.

In other words, the server desires that the user agent engage in reactive negotiation to select the most appropriate representation(s) for its needs<sup><a href="#ref-1">1</a></sup>.

If the server has a preferred choice, the server SHOULD generate a Location header field containing a preferred choice&#39;s URI reference. The user agent MAY use the Location field value for automatic redirection.

For request methods other than HEAD, the server SHOULD generate a payload in the 300 response containing a list of representation metadata and URI reference(s) from which the user or user agent can choose the one most preferred. The user agent MAY make a selection from that list automatically if it understands the provided media type. A specific format for automatic selection is not defined by this specification because HTTP tries to remain orthogonal to the definition of its payloads. In practice, the representation is provided in some easily parsed format believed to be acceptable to the user agent, as determined by shared design or content negotiation, or in some commonly accepted hypertext format.

A 300 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-2">2</a></sup>.

Note: The original proposal for the 300 status code defined the URI header field as providing a list of alternative representations, such that it would be usable for 200, 300, and 406 responses and be transferred in responses to the HEAD method. However, lack of deployment and disagreement over syntax led to both URI and Alternates (a subsequent proposal) being dropped from this specification. It is possible to communicate the list using a set of Link header fields<sup><a href="#ref-3">3</a></sup>, each with a relationship of &quot;alternate&quot;, though deployment is a chicken-and-egg problem.
`,
	},
	{
		Code: 301,
		Name: "Moved Permanently",
		Description: `

The target resource has been assigned a new permanent URI and any future references to this resource ought to use one of the enclosed URIs.

Clients with link-editing capabilities ought to automatically re-link references to the effective request URI to one or more of the new references sent by the server, where possible.

The server SHOULD generate a Location header field in the response containing a preferred URI reference for the new permanent URI. The user agent MAY use the Location field value for automatic redirection. The server&#39;s response payload usually contains a short hypertext note with a hyperlink to the new URI(s).

Note: For historical reasons, a user agent MAY change the request method from POST to GET for the subsequent request. If this behavior is undesired, the <a href="/307">307 Temporary Redirect</a> status code can be used instead.

A 301 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 302,
		Name: "Found",
		Description: `

The target resource resides temporarily under a different URI. Since the redirection might be altered on occasion, the client ought to continue to use the effective request URI for future requests.

The server SHOULD generate a Location header field in the response containing a URI reference for the different URI. The user agent MAY use the Location field value for automatic redirection. The server&#39;s response payload usually contains a short hypertext note with a hyperlink to the different URI(s).

Note: For historical reasons, a user agent MAY change the request method from POST to GET for the subsequent request. If this behavior is undesired, the <a href="/307">307 Temporary Redirect</a> status code can be used instead.
`,
	},
	{
		Code: 303,
		Name: "See Other",
		Description: `

The server is redirecting the user agent to a different resource, as indicated by a URI in the Location header field, which is intended to provide an indirect response to the original request.

A user agent can perform a retrieval request targeting that URI (a GET or HEAD request if using HTTP), which might also be redirected, and present the eventual result as an answer to the original request. Note that the new URI in the Location header field is not considered equivalent to the effective request URI.

This status code is applicable to any HTTP method. It is primarily used to allow the output of a POST action to redirect the user agent to a selected resource, since doing so provides the information corresponding to the POST response in a form that can be separately identified, bookmarked, and cached, independent of the original request.

A 303 response to a GET request indicates that the origin server does not have a representation of the target resource that can be transferred by the server over HTTP. However, the Location field value refers to a resource that is descriptive of the target resource, such that making a retrieval request on that other resource might result in a representation that is useful to recipients without implying that it represents the original target resource. Note that answers to the questions of what can be represented, what representations are adequate, and what might be a useful description are outside the scope of HTTP.

Except for responses to a HEAD request, the representation of a 303 response ought to contain a short hypertext note with a hyperlink to the same URI reference provided in the Location header field.
`,
	},
	{
		Code: 304,
		Name: "Not Modified",
		Description: `

A conditional GET or HEAD request has been received and would have resulted in a <a href="/200">200 OK</a> response if it were not for the fact that the condition evaluated to false.

In other words, there is no need for the server to transfer a representation of the target resource because the request indicates that the client, which made the request conditional, already has a valid representation; the server is therefore redirecting the client to make use of that stored representation as if it were the payload of a <a href="/200">200 OK</a> response.

The server generating a 304 response MUST generate any of the following header fields that would have been sent in a <a href="/200">200 OK</a> response to the same request: Cache-Control, Content-Location, Date, ETag, Expires, and Vary.

Since the goal of a 304 response is to minimize information transfer when the recipient already has one or more cached representations, a sender SHOULD NOT generate representation metadata other than the above listed fields unless said metadata exists for the purpose of guiding cache updates (e.g., Last-Modified might be useful if the response does not have an ETag field).

Requirements on a cache that receives a 304 response are defined in <a href="https://tools.ietf.org/html/rfc7234#section-4.3.4">Section 4.3.4 of RFC7234</a>. If the conditional request originated with an outbound client, such as a user agent with its own cache sending a conditional GET to a shared proxy, then the proxy SHOULD forward the 304 response to that client.

A 304 response cannot contain a message-body; it is always terminated by the first empty line after the header fields.
`,
	},
	{
		Code: 305,
		Name: "Use Proxy",
		Description: `

Defined in a previous version of this specification and is now deprecated, due to security concerns regarding in-band configuration of a proxy.
`,
	},
	{
		Code: 307,
		Name: "Temporary Redirect",
		Description: `

The target resource resides temporarily under a different URI and the user agent MUST NOT change the request method if it performs an automatic redirection to that URI.

Since the redirection can change over time, the client ought to continue using the original effective request URI for future requests.

The server SHOULD generate a Location header field in the response containing a URI reference for the different URI. The user agent MAY use the Location field value for automatic redirection. The server&#39;s response payload usually contains a short hypertext note with a hyperlink to the different URI(s).

Note: This status code is similar to 302 Found, except that it does not allow changing the request method from POST to GET. This specification defines no equivalent counterpart for <a href="/301">301 Moved Permanently</a> (<a href="https://tools.ietf.org/html/rfc7238">RFC7238</a>, however, proposes defining the status code <a href="/308">308 Permanent Redirect</a> for this purpose).
`,
	},
	{
		Code: 308,
		Name: "Permanent Redirect",
		Description: `

The target resource has been assigned a new permanent URI and any future references to this resource ought to use one of the enclosed URIs.

Clients with link editing capabilities ought to automatically re-link references to the effective request URI<sup><a href="#ref-1">1</a></sup> to one or more of the new references sent by the server, where possible.

The server SHOULD generate a Location header field<sup><a href="#ref-2">2</a></sup> in the response containing a preferred URI reference for the new permanent URI. The user agent MAY use the Location field value for automatic redirection. The server&#39;s response payload usually contains a short hypertext note with a hyperlink to the new URI(s).

A 308 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-3">3</a></sup>.

Note: This status code is similar to <a href="/301">301 Moved Permanently</a>, except that it does not allow changing the request method from POST to GET.
`,
	},
	{
		Code: 400,
		Name: "Bad Request",
		Description: `

The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
`,
	},
	{
		Code: 401,
		Name: "Unauthorized",
		Description: `

The request has not been applied because it lacks valid authentication credentials for the target resource.

The server generating a 401 response MUST send a WWW-Authenticate header field<sup><a href="#ref-1">1</a></sup> containing at least one challenge applicable to the target resource.

If the request included authentication credentials, then the 401 response indicates that authorization has been refused for those credentials. The user agent MAY repeat the request with a new or replaced Authorization header field<sup><a href="#ref-2">2</a></sup>. If the 401 response contains the same challenge as the prior response, and the user agent has already attempted authentication at least once, then the user agent SHOULD present the enclosed representation to the user, since it usually contains relevant diagnostic information.
`,
	},
	{
		Code: 402,
		Name: "Payment Required",
		Description: `

Reserved for future use.
`,
	},
	{
		Code: 403,
		Name: "Forbidden",
		Description: `

The server understood the request but refuses to authorize it.

A server that wishes to make public why the request has been forbidden can describe that reason in the response payload (if any).

If authentication credentials were provided in the request, the server considers them insufficient to grant access. The client SHOULD NOT automatically repeat the request with the same credentials. The client MAY repeat the request with new or different credentials. However, a request might be forbidden for reasons unrelated to the credentials.

An origin server that wishes to &quot;hide&quot; the current existence of a forbidden target resource MAY instead respond with a status code of <a href="/404">404 Not Found</a>.
`,
	},
	{
		Code: 404,
		Name: "Not Found",
		Description: `

The origin server did not find a current representation for the target resource or is not willing to disclose that one exists.

A 404 status code does not indicate whether this lack of representation is temporary or permanent; the <a href="/410">410 Gone</a> status code is preferred over 404 if the origin server knows, presumably through some configurable means, that the condition is likely to be permanent.

A 404 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 405,
		Name: "Method Not Allowed",
		Description: `

The method received in the request-line is known by the origin server but not supported by the target resource.

The origin server MUST generate an Allow header field in a 405 response containing a list of the target resource&#39;s currently supported methods.

A 405 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 406,
		Name: "Not Acceptable",
		Description: `

The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request<sup><a href="#ref-1">1</a></sup>, and the server is unwilling to supply a default representation.

The server SHOULD generate a payload containing a list of available representation characteristics and corresponding resource identifiers from which the user or user agent can choose the one most appropriate. A user agent MAY automatically select the most appropriate choice from that list. However, this specification does not define any standard for such automatic selection, as described in <a href="https://tools.ietf.org/html/rfc7231#section-6.4.1">RFC7231 Section 6.4.1</a>.
`,
	},
	{
		Code: 407,
		Name: "Proxy Authentication Required",
		Description: `

Similar to <a href="/401">401 Unauthorized</a>, but it indicates that the client needs to authenticate itself in order to use a proxy.

The proxy MUST send a Proxy-Authenticate header field<sup><a href="#ref-1">1</a></sup> containing a challenge applicable to that proxy for the target resource. The client MAY repeat the request with a new or replaced Proxy-Authorization header field<sup><a href="#ref-2">2</a></sup>.
`,
	},
	{
		Code: 408,
		Name: "Request Timeout",
		Description: `

The server did not receive a complete request message within the time that it was prepared to wait.

A server SHOULD send the &quot;close&quot; connection option<sup><a href="#ref-1">1</a></sup> in the response, since 408 implies that the server has decided to close the connection rather than continue waiting. If the client has an outstanding request in transit, the client MAY repeat that request on a new connection.
`,
	},
	{
		Code: 409,
		Name: "Conflict",
		Description: `

The request could not be completed due to a conflict with the current state of the target resource. This code is used in situations where the user might be able to resolve the conflict and resubmit the request.

The server SHOULD generate a payload that includes enough information for a user to recognize the source of the conflict.

Conflicts are most likely to occur in response to a PUT request. For example, if versioning were being used and the representation being PUT included changes to a resource that conflict with those made by an earlier (third-party) request, the origin server might use a 409 response to indicate that it can&#39;t complete the request. In this case, the response representation would likely contain information useful for merging the differences based on the revision history.
`,
	},
	{
		Code: 410,
		Name: "Gone",
		Description: `

The target resource is no longer available at the origin server and that this condition is likely to be permanent.

If the origin server does not know, or has no facility to determine, whether or not the condition is permanent, the status code <a href="/404">404 Not Found</a> ought to be used instead.

The 410 response is primarily intended to assist the task of web maintenance by notifying the recipient that the resource is intentionally unavailable and that the server owners desire that remote links to that resource be removed. Such an event is common for limited-time, promotional services and for resources belonging to individuals no longer associated with the origin server&#39;s site. It is not necessary to mark all permanently unavailable resources as &quot;gone&quot; or to keep the mark for any length of time -- that is left to the discretion of the server owner.

A 410 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 411,
		Name: "Length Required",
		Description: `

The server refuses to accept the request without a defined Content-Length<sup><a href="#ref-1">1</a></sup>.

The client MAY repeat the request if it adds a valid Content-Length header field containing the length of the message body in the request message.
`,
	},
	{
		Code: 412,
		Name: "Precondition Failed",
		Description: `

One or more conditions given in the request header fields evaluated to false when tested on the server.

This response code allows the client to place preconditions on the current resource state (its current representations and metadata) and, thus, prevent the request method from being applied if the target resource is in an unexpected state.
`,
	},
	{
		Code: 413,
		Name: "Payload Too Large",
		Description: `

The server is refusing to process a request because the request payload is larger than the server is willing or able to process.

The server MAY close the connection to prevent the client from continuing the request.

If the condition is temporary, the server SHOULD generate a Retry-After header field to indicate that it is temporary and after what time the client MAY try again.
`,
	},
	{
		Code: 414,
		Name: "Request-URI Too Long",
		Description: `

The server is refusing to service the request because the request-target<sup><a href="#ref-1">1</a></sup> is longer than the server is willing to interpret.

This rare condition is only likely to occur when a client has improperly converted a POST request to a GET request with long query information, when the client has descended into a &quot;black hole&quot; of redirection (e.g., a redirected URI prefix that points to a suffix of itself) or when the server is under attack by a client attempting to exploit potential security holes.

A 414 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls <sup><a href="#ref-2">2</a></sup>.
`,
	},
	{
		Code: 415,
		Name: "Unsupported Media Type",
		Description: `

The origin server is refusing to service the request because the payload is in a format not supported by this method on the target resource.

The format problem might be due to the request&#39;s indicated Content-Type or Content-Encoding, or as a result of inspecting the data directly.
`,
	},
	{
		Code: 416,
		Name: "Requested Range Not Satisfiable",
		Description: `

None of the ranges in the request&#39;s Range header field<sup><a href="#ref-1">1</a></sup> overlap the current extent of the selected resource or that the set of ranges requested has been rejected due to invalid ranges or an excessive request of small or overlapping ranges.

For byte ranges, failing to overlap the current extent means that the first-byte-pos of all of the byte-range-spec values were greater than the current length of the selected representation. When this status code is generated in response to a byte-range request, the sender SHOULD generate a Content-Range header field specifying the current length of the selected representation<sup><a href="#ref-2">2</a></sup>.

For example:
<pre><code>HTTP/1.1 416 Range Not Satisfiable
Date: Fri, 20 Jan 2012 15:41:54 GMT
Content-Range: bytes */47022
</code></pre>
Note: Because servers are free to ignore Range, many implementations will simply respond with the entire selected representation in a <a href="/200">200 OK</a> response. That is partly because most clients are prepared to receive a <a href="/200">200 OK</a> to complete the task (albeit less efficiently) and partly because clients might not stop making an invalid partial request until they have received a complete representation. Thus, clients cannot depend on receiving a 416 Range Not Satisfiable response even when it is most appropriate.
`,
	},
	{
		Code: 417,
		Name: "Expectation Failed",
		Description: `

The expectation given in the request&#39;s Expect header field<sup><a href="#ref-1">1</a></sup> could not be met by at least one of the inbound servers.
`,
	},
	{
		Code: 418,
		Name: "> I'm a teapot",
		Description: `

Any attempt to brew coffee with a teapot should result in the error code &quot;418 I&#39;m a teapot&quot;. The resulting entity body MAY be short and stout.
`,
	},
	{
		Code: 421,
		Name: "Misdirected Request",
		Description: `

The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.

Clients receiving a 421 Misdirected Request response from a server MAY retry the request -- whether the request method is idempotent or not -- over a different connection. This is possible if a connection is reused<sup><a href="#ref-1">1</a></sup>or if an alternative service is selected <a href="https://tools.ietf.org/html/rfc7540#ref-ALT-SVC">ALT-SVC</a>.

This status code MUST NOT be generated by proxies.

A 421 response is cacheable by default, i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-2">2</a></sup>.
`,
	},
	{
		Code: 422,
		Name: "Unprocessable Entity",
		Description: `

The server understands the content type of the request entity (hence a <a href="/415">415 Unsupported Media Type</a> status code is inappropriate), and the syntax of the request entity is correct (thus a <a href="/400">400 Bad Request</a> status code is inappropriate) but was unable to process the contained instructions.

For example, this error condition may occur if an XML request body contains well-formed (i.e., syntactically correct), but semantically erroneous, XML instructions.
`,
	},
	{
		Code: 423,
		Name: "Locked",
		Description: `

The source or destination resource of a method is locked.

This response SHOULD contain an appropriate precondition or postcondition code, such as &#39;lock-token-submitted&#39; or &#39;no-conflicting-lock&#39;.
`,
	},
	{
		Code: 424,
		Name: "Failed Dependency",
		Description: `

The method could not be performed on the resource because the requested action depended on another action and that action failed.

For example, if a command in a PROPPATCH method fails, then, at minimum, the rest of the commands will also fail with 424 Failed Dependency.
`,
	},
	{
		Code: 426,
		Name: "Upgrade Required",
		Description: `

The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol.

The server MUST send an Upgrade header field in a 426 response to indicate the required protocol(s)<sup><a href="#ref-1">1</a></sup>

Example:
<pre><code>HTTP/1.1 426 Upgrade Required
Upgrade: HTTP/3.0
Connection: Upgrade
Content-Length: 53
Content-Type: text/plain

This service requires use of the HTTP/3.0 protocol.
`,
	},
	{
		Code: 428,
		Name: "Precondition Required",
		Description: `

The origin server requires the request to be conditional.

Its typical use is to avoid the &quot;lost update&quot; problem, where a client GETs a resource&#39;s state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict. By requiring requests to be conditional, the server can assure that clients are working with the correct copies.

Responses using this status code SHOULD explain how to resubmit the request successfully. For example:
<pre><code>HTTP/1.1 428 Precondition Required
Content-Type: text/html

&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;Precondition Required&lt;/title&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;h1&gt;Precondition Required&lt;/h1&gt;
    &lt;p&gt;This request is required to be conditional; try using &quot;If-Match&quot;.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;
</code></pre>
Responses with the 428 status code MUST NOT be stored by a cache.
`,
	},
	{
		Code: 429,
		Name: "Too Many Requests",
		Description: `

The user has sent too many requests in a given amount of time (&quot;rate limiting&quot;).

The response representations SHOULD include details explaining the condition, and MAY include a Retry-After header indicating how long to wait before making a new request.

For example:
<pre><code>HTTP/1.1 429 Too Many Requests
Content-Type: text/html
Retry-After: 3600

&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;Too Many Requests&lt;/title&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;h1&gt;Too Many Requests&lt;/h1&gt;
    &lt;p&gt;I only allow 50 requests per hour to this Web site per
    logged in user. Try again soon.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;
</code></pre>
Note that this specification does not define how the origin server identifies the user, nor how it counts requests. For example, an origin server that is limiting request rates can do so based upon counts of requests on a per-resource basis, across the entire server, or even among a set of servers. Likewise, it might identify the user by its authentication credentials, or a stateful cookie.

Responses with the 429 status code MUST NOT be stored by a cache.
`,
	},
	{
		Code: 431,
		Name: "Request Header Fields Too Large",
		Description: `

The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields.

It can be used both when the set of request header fields in total is too large, and when a single header field is at fault.  In the latter case, the response representation SHOULD specify which header field was too large.

For example:
<pre><code>HTTP/1.1 431 Request Header Fields Too Large
Content-Type: text/html

&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;Request Header Fields Too Large&lt;/title&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;h1&gt;Request Header Fields Too Large&lt;/h1&gt;
    &lt;p&gt;The &quot;Example&quot; header was too large.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;
</code></pre>
Responses with the 431 status code MUST NOT be stored by a cache.
`,
	},
	{
		Code: 444,
		Name: "Connection Closed Without Response",
		Description: `

A non-standard status code used to instruct <a href="https://nginx.org">nginx</a> to close the connection without sending a response to the client, most commonly used to deny malicious or malformed requests.

This status code is not seen by the client, it only appears in nginx log files.
`,
	},
	{
		Code: 451,
		Name: "Unavailable For Legal Reasons",
		Description: `

The server is denying access to the resource as a consequence of a legal demand.

The server in question might not be an origin server. This type of legal demand typically most directly affects the operations of ISPs and search engines.

Responses using this status code SHOULD include an explanation, in the response body, of the details of the legal demand: the party making it, the applicable legislation or regulation, and what classes of person and resource it applies to. For example:
<pre><code>HTTP/1.1 451 Unavailable For Legal Reasons
Link: &lt;https://spqr.example.org/legislatione&gt;; rel=&quot;blocked-by&quot;
Content-Type: text/html

&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;Unavailable For Legal Reasons&lt;/title&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;h1&gt;Unavailable For Legal Reasons&lt;/h1&gt;
    &lt;p&gt;This request may not be serviced in the Roman Province
    of Judea due to the Lex Julia Majestatis, which disallows
    access to resources hosted on servers deemed to be
    operated by the People&#39;s Front of Judea.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;
</code></pre>
The use of the 451 status code implies neither the existence nor non- existence of the resource named in the request. That is to say, it is possible that if the legal demands were removed, a request for the resource still might not succeed.

Note that in many cases clients can still access the denied resource by using technical countermeasures such as a VPN or the Tor network.

A 451 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls; see <a href="https://tools.ietf.org/html/rfc7234">RFC7234</a>.
`,
	},
	{
		Code: 499,
		Name: "Client Closed Request",
		Description: `

A non-standard status code introduced by <a href="https://nginx.org">nginx</a> for the case when a client closes the connection while nginx is processing the request.
`,
	},
	{
		Code: 500,
		Name: "Internal Server Error",
		Description: `

The server encountered an unexpected condition that prevented it from fulfilling the request.
`,
	},
	{
		Code: 501,
		Name: "Not Implemented",
		Description: `

The server does not support the functionality required to fulfill the request.

This is the appropriate response when the server does not recognize the request method and is not capable of supporting it for any resource.

A 501 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls<sup><a href="#ref-1">1</a></sup>.
`,
	},
	{
		Code: 502,
		Name: "Bad Gateway",
		Description: `

The server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while attempting to fulfill the request.
`,
	},
	{
		Code: 503,
		Name: "Service Unavailable",
		Description: `

The server is currently unable to handle the request due to a temporary overload or scheduled maintenance, which will likely be alleviated after some delay.

The server MAY send a Retry-After header field<sup><a href="#ref-1">1</a></sup> to suggest an appropriate amount of time for the client to wait before retrying the request.

Note: The existence of the 503 status code does not imply that a server has to use it when becoming overloaded. Some servers might simply refuse the connection.
`,
	},
	{
		Code: 504,
		Name: "Gateway Timeout",
		Description: `

The server, while acting as a gateway or proxy, did not receive a timely response from an upstream server it needed to access in order to complete the request.
`,
	},
	{
		Code: 505,
		Name: "HTTP Version Not Supported",
		Description: `

The server does not support, or refuses to support, the major version of HTTP that was used in the request message.

The server is indicating that it is unable or unwilling to complete the request using the same major version as the client, as described in <a href="https://tools.ietf.org/html/rfc7230#section-2.6">Section 2.6 of RFC7230</a>, other than with this error message. The server SHOULD generate a representation for the 505 response that describes why that version is not supported and what other protocols are supported by that server.
`,
	},
	{
		Code: 506,
		Name: "Variant Also Negotiates",
		Description: `

The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.
`,
	},
	{
		Code: 507,
		Name: "Insufficient Storage",
		Description: `

The method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request.

This condition is considered to be temporary. If the request that received this status code was the result of a user action, the request MUST NOT be repeated until it is requested by a separate user action.
`,
	},
	{
		Code: 508,
		Name: "Loop Detected",
		Description: `

The server terminated an operation because it encountered an infinite loop while processing a request with &quot;Depth: infinity&quot;. This status indicates that the entire operation failed.
`,
	},
	{
		Code: 510,
		Name: "Not Extended",
		Description: `

The policy for accessing the resource has not been met in the request. The server should send back all the information necessary for the client to issue an extended request.

It is outside the scope of this specification to specify how the extensions inform the client.

If the 510 response contains information about extensions that were not present in the initial request then the client MAY repeat the request if it has reason to believe it can fulfill the extension policy by modifying the request according to the information provided in the 510 response. Otherwise the client MAY present any entity included in the 510 response to the user, since that entity may include relevant diagnostic information.
`,
	},
	{
		Code: 511,
		Name: "Network Authentication Required",
		Description: `

The client needs to authenticate to gain network access.

The response representation SHOULD contain a link to a resource that allows the user to submit credentials (e.g., with an HTML form).

Note that the 511 response SHOULD NOT contain a challenge or the login interface itself, because browsers would show the login interface as being associated with the originally requested URL, which may cause confusion.

The 511 status SHOULD NOT be generated by origin servers; it is intended for use by intercepting proxies that are interposed as a means of controlling access to the network.

Responses with the 511 status code MUST NOT be stored by a cache.

The 511 status code is designed to mitigate problems caused by &quot;captive portals&quot; to software (especially non-browser agents) that is expecting a response from the server that a request was made to, not the intervening network infrastructure. It is not intended to encourage deployment of captive portals -- only to limit the damage caused by them.

A network operator wishing to require some authentication, acceptance of terms, or other user interaction before granting access usually does so by identifying clients who have not done so (&quot;unknown clients&quot;) using their Media Access Control (MAC) addresses.

Unknown clients then have all traffic blocked, except for that on TCP port 80, which is sent to an HTTP server (the &quot;login server&quot;) dedicated to &quot;logging in&quot; unknown clients, and of course traffic to the login server itself.

For example, a user agent might connect to a network and make the following HTTP request on TCP port 80:
<pre><code>GET /index.htm HTTP/1.1
Host: www.example.com
</code></pre>
Upon receiving such a request, the login server would generate a 511 response:
<pre><code>HTTP/1.1 511 Network Authentication Required
Content-Type: text/html

&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;Network Authentication Required&lt;/title&gt;
    &lt;meta http-equiv=&quot;refresh&quot; content=&quot;0; url=https://login.example.net/&quot;&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;p&gt;You need to &lt;a href=&quot;https://login.example.net/&quot;&gt;
    authenticate with the local network&lt;/a&gt; in order to gain
    access.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;
</code></pre>
Here, the 511 status code assures that non-browser clients will not interpret the response as being from the origin server, and the META HTML element redirects the user agent to the login server.
`,
	},
	{
		Code: 599,
		Name: "Network Connect Timeout Error",
		Description: `

This status code is not specified in any RFCs, but is used by some HTTP proxies to signal a network connect timeout behind the proxy to a client in front of the proxy.
`,
	},
}