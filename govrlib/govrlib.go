package main

import "C"
import (
	"fmt"

	gl "github.com/go-gl/gl/v3.1/gles2"
)

var (
	buffer              uint32
	vao                 uint32
	program             uint32
	vertexLoc           int32
	viewMatrixLoc       int32
	projectionMatrixLoc int32
)

const vssrc = `
#version 300 es
#define NUM_VIEWS 2
#extension GL_OVR_multiview2 : enable
layout(num_views=NUM_VIEWS) in;
#define VIEW_ID gl_ViewID_OVR

in vec3 vertex;
uniform mat4 ViewMatrix[NUM_VIEWS];
uniform mat4 ProjectionMatrix[NUM_VIEWS];

void main()
{
	gl_Position = ProjectionMatrix[VIEW_ID] * ViewMatrix[VIEW_ID] * vec4( vertex, 1.0 );
}
` + "\x00"

const fssrc = `
#version 300 es

out lowp vec4 outColor;
void main()
{
	outColor = vec4(0.0, 1.0, 0.0, 1.0);
}
` + "\x00"

func main() {}

//export Init
func Init() {
	gl.Init()

	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	data := [9]float32{
		0.0, 0.0, -2.0,
		1.0, 0.0, -2.0,
		0.5, 1.0, -2.0,
	}
	gl.BufferData(gl.ARRAY_BUFFER, 9*4, gl.Ptr(data[:]), gl.STATIC_DRAW)

	var vs, fs uint32
	var status int32

	var buf [65536]byte
	var length int32

	vs = gl.CreateShader(gl.VERTEX_SHADER)
	cstrs1, free1 := gl.Strs(vssrc)
	defer free1()
	gl.ShaderSource(vs, 1, cstrs1, nil)
	gl.CompileShader(vs)
	gl.GetShaderiv(vs, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE {
		log("vertex shader compilation failed")
		gl.GetShaderInfoLog(vs, 65535, &length, &buf[0])
		log(fmt.Sprintf("%s", buf))
		return
	}

	fs = gl.CreateShader(gl.FRAGMENT_SHADER)
	cstrs2, free2 := gl.Strs(fssrc)
	defer free2()
	gl.ShaderSource(fs, 1, cstrs2, nil)
	gl.CompileShader(fs)
	gl.GetShaderiv(fs, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE {
		log("fragment shader compilation failed")
		gl.GetShaderInfoLog(vs, 65535, &length, &buf[0])
		log(fmt.Sprintf("%s", buf))
		return
	}

	program = gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)

	gl.LinkProgram(program)
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status != gl.TRUE {
		log("program link failed")
		gl.GetProgramInfoLog(program, 65535, &length, &buf[0])
		log(fmt.Sprintf("%s", buf))
		return
	}

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.UseProgram(program)
	vertexLoc = gl.GetAttribLocation(program, gl.Str("vertex\x00"))
	viewMatrixLoc = gl.GetUniformLocation(program, gl.Str("ViewMatrix\x00"))
	projectionMatrixLoc = gl.GetUniformLocation(program, gl.Str("ProjectionMatrix\x00"))
}

//export Run
func Run() {
}

//export Render
func Render(eye, projection *float32) {
	gl.UseProgram(program)

	gl.UniformMatrix4fv(viewMatrixLoc, 2, false, eye)
	gl.UniformMatrix4fv(projectionMatrixLoc, 2, false, projection)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.EnableVertexAttribArray(uint32(vertexLoc))
	gl.VertexAttribPointer(uint32(vertexLoc), 3, gl.FLOAT, false, 12, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
