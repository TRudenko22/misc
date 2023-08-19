// Turners cube

module neocube(size, center) {
    difference() {
        cube(size, center);
        
        cylinder(r = size/2.2,h=size + 2, center=true);
        rotate([90, 0, 0]) cylinder(r = size/2.2,h=size + 2, center=true);
        rotate([0, 90, 0]) cylinder(r = size/2.2,h=size + 2, center=true);


    }
}
