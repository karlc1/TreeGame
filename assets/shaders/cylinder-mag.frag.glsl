#version 330 core

const float radius=2.;
const float depth=radius/2.;

out vec4 fragColor;
uniform sampler2D uTexture;
uniform vec4 uTexBounds;

// === main loop ===
void main() {

	// TODO: uTexBounds.xy?

	vec2 uv = gl_FragCoord.xy / uTexBounds.zw;
	vec2 center = vec2(uTexBounds.z / 2, uTexBounds.w / 2);
	//vec2 center = vec2(uTexBounds.xy / 2.0);
    
	float ax = ((uv.x - center.x) * (uv.x - center.x)) / (0.2*0.2); 
    
	float dx = 0.0 + (-depth/radius)*ax + (depth/(radius*radius))*ax*ax;
    
    	float f =  (ax + dx );
    
	if (ax > radius) f = ax;
    
    	vec2 magnifierArea = center + (uv-center)*f/ax;
    
    	fragColor = vec4(texture( uTexture, vec2(1,-1) * magnifierArea ).rgb, 0.);  
}

