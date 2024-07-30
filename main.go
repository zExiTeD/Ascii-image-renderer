package main

import (
    "image"
    "image/color"
    "image/png"
    "os"

    rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
    const screenWidth, screenHeight = 800, 600

    // Initialize the window
    rl.InitWindow(screenWidth, screenHeight, "Open Image Pixel by Pixel with Zoom")
    defer rl.CloseWindow()

    // Load the image using the png package
    imagePath := "path/to/your/image.png"
    file, err := os.Open(imagePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    img, err := png.Decode(file)
    if err != nil {
        panic(err)
    }

    // Create a new Raylib image
    width, height := img.Bounds().Dx(), img.Bounds().Dy()
    rayImage := rl.GenImageColor(int32(width), int32(height), rl.RayWhite)

    // Access and manipulate pixels
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            pixel := img.At(x, y).(color.RGBA)
            newColor := manipulatePixelColor(pixel)
            rl.ImageDrawPixel(rayImage, int32(x), int32(y), newColor)
        }
    }

    // Create a texture from the Raylib image
    texture := rl.LoadTextureFromImage(rayImage)
    defer rl.UnloadTexture(texture)

    // Initial zoom level and position
    zoomLevel := 1.0
    zoomStep := 0.1
    position := rl.NewVector2(0, 0)

    // Main game loop
    for !rl.WindowShouldClose() {
        // Handle zoom
        if rl.IsKeyPressed(rl.KeyUp) {
            zoomLevel += zoomStep
        }
        if rl.IsKeyPressed(rl.KeyDown) {
            zoomLevel -= zoomStep
            if zoomLevel < zoomStep {
                zoomLevel = zoomStep
            }
        }

        // Get mouse wheel movement for zoom
        wheel := rl.GetMouseWheelMove()
        if wheel > 0 {
            zoomLevel += zoomStep
        } else if wheel < 0 {
            zoomLevel -= zoomStep
            if zoomLevel < zoomStep {
                zoomLevel = zoomStep
            }
        }

        // Update position based on mouse dragging
        if rl.IsMouseButtonDown(rl.MouseLeftButton) {
            mouseDelta := rl.GetMouseDelta()
            position.X += mouseDelta.X
            position.Y += mouseDelta.Y
        }

        // Begin drawing
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)

        // Draw the manipulated image with zoom
        rl.DrawTextureEx(texture, position, 0, float32(zoomLevel), rl.White)

        
        rl.EndDrawing()
    }
}

// Manipulate the pixel color (example: invert colors)
func manipulatePixelColor(c color.RGBA) rl.Color {
    return rl.Color{255 - c.R, 255 - c.G, 255 - c.B, c.A}
}


