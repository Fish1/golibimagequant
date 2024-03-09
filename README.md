# golibimagequant

golibimagequant is binding library for [libimagequant](https://github.com/ImageOptim/libimagequant)

## Supported Features

### attr creation
- [x] liq_attr_create
- [ ] liq_attr_create_with_allocator
- [x] liq_attr_copy
- [x] liq_attr_destroy

### histogram creation
- [ ] liq_histogram_create
- [ ] liq_histogram_add_image
- [ ] liq_histogram_destroy

### histogram controls
- [ ] liq_histogram_add_colors
- [ ] liq_histogram_add_fixed_color

### quality controls
- [x] liq_set_max_colors
- [x] liq_get_max_colors
- [x] liq_set_speed
- [x] liq_get_speed
- [x] liq_set_min_posterization
- [x] liq_get_min_posterization
- [x] liq_set_quality
- [x] liq_get_min_quality
- [x] liq_get_max_quality
- [ ] liq_set_last_index_transparent

### logging
- [ ] liq_set_log_callback
- [ ] liq_set_log_flush_callback
- [ ] liq_attr_set_progress_callback
- [ ] liq_result_set_progress_callback

### image creation
- [ ] liq_image_create_rgba_rows
- [x] liq_image_create_rgba
- [ ] liq_image_create_custom
- [x] liq_image_destroy

### image controls
- [ ] liq_image_set_memory_ownership
- [ ] liq_image_set_background
- [ ] liq_image_set_importance_map
- [x] liq_image_add_fixed_color
- [ ] liq_image_get_width
- [ ] liq_image_get_height

### quantization
- [ ] liq_histogram_quantize
- [x] liq_image_quantize
- [x] liq_result_destroy

### quantization controls
- [ ] liq_set_dithering_level
- [ ] liq_set_output_gamma
- [ ] liq_get_output_gama

### quantization results
- [x] liq_get_palette
- [x] liq_write_remapped_image
- [ ] liq_write_remapped_image_rows

### quality control
- [ ] liq_get_quantization_error
- [ ] liq_get_quantization_quality
- [ ] liq_get_remapping_error
- [ ] liq_get_remapping_quality

### other
- [x] liq_version
