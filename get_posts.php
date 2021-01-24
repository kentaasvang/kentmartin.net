<?php

namespace Kentmartin;


function getPosts ( $folder ) {
    $files = [];

    foreach ( $folder as $file ) 
    {
        // skip if it is a dotfile
        if ( $file->isDot() )
            continue;

        array_push ( $files, $file->getFilename() );
    }
    
    return $files;
}
