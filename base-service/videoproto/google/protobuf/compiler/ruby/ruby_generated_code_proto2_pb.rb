# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ruby_generated_code_proto2.proto

require 'google/protobuf'

require 'ruby_generated_code_proto2_import_pb'


descriptor_data = "\n ruby_generated_code_proto2.proto\x12\x05\x41.B.C\x1a\'ruby_generated_code_proto2_import.proto\"\x96\x0b\n\x0bTestMessage\x12\x19\n\x0eoptional_int32\x18\x01 \x01(\x05:\x01\x31\x12\x19\n\x0eoptional_int64\x18\x02 \x01(\x03:\x01\x32\x12\x1a\n\x0foptional_uint32\x18\x03 \x01(\r:\x01\x33\x12\x1a\n\x0foptional_uint64\x18\x04 \x01(\x04:\x01\x34\x12\x1b\n\roptional_bool\x18\x05 \x01(\x08:\x04true\x12\x1a\n\x0foptional_double\x18\x06 \x01(\x01:\x01\x36\x12\x19\n\x0eoptional_float\x18\x07 \x01(\x02:\x01\x37\x12$\n\x0foptional_string\x18\x08 \x01(\t:\x0b\x64\x65\x66\x61ult str\x12*\n\x0eoptional_bytes\x18\t \x01(\x0c:\x12\\000\\001\\002@fubar\x12)\n\roptional_enum\x18\n \x01(\x0e\x32\x0f.A.B.C.TestEnum:\x01\x41\x12(\n\x0coptional_msg\x18\x0b \x01(\x0b\x32\x12.A.B.C.TestMessage\x12>\n\x1aoptional_proto2_submessage\x18\x0c \x01(\x0b\x32\x1a.A.B.C.TestImportedMessage\x12\x16\n\x0erepeated_int32\x18\x15 \x03(\x05\x12\x16\n\x0erepeated_int64\x18\x16 \x03(\x03\x12\x17\n\x0frepeated_uint32\x18\x17 \x03(\r\x12\x17\n\x0frepeated_uint64\x18\x18 \x03(\x04\x12\x15\n\rrepeated_bool\x18\x19 \x03(\x08\x12\x17\n\x0frepeated_double\x18\x1a \x03(\x01\x12\x16\n\x0erepeated_float\x18\x1b \x03(\x02\x12\x17\n\x0frepeated_string\x18\x1c \x03(\t\x12\x16\n\x0erepeated_bytes\x18\x1d \x03(\x0c\x12&\n\rrepeated_enum\x18\x1e \x03(\x0e\x32\x0f.A.B.C.TestEnum\x12(\n\x0crepeated_msg\x18\x1f \x03(\x0b\x32\x12.A.B.C.TestMessage\x12\x16\n\x0erequired_int32\x18) \x02(\x05\x12\x16\n\x0erequired_int64\x18* \x02(\x03\x12\x17\n\x0frequired_uint32\x18+ \x02(\r\x12\x17\n\x0frequired_uint64\x18, \x02(\x04\x12\x15\n\rrequired_bool\x18- \x02(\x08\x12\x17\n\x0frequired_double\x18. \x02(\x01\x12\x16\n\x0erequired_float\x18/ \x02(\x02\x12\x17\n\x0frequired_string\x18\x30 \x02(\t\x12\x16\n\x0erequired_bytes\x18\x31 \x02(\x0c\x12&\n\rrequired_enum\x18\x32 \x02(\x0e\x32\x0f.A.B.C.TestEnum\x12(\n\x0crequired_msg\x18\x33 \x02(\x0b\x32\x12.A.B.C.TestMessage\x12\x15\n\x0boneof_int32\x18= \x01(\x05H\x00\x12\x15\n\x0boneof_int64\x18> \x01(\x03H\x00\x12\x16\n\x0coneof_uint32\x18? \x01(\rH\x00\x12\x16\n\x0coneof_uint64\x18@ \x01(\x04H\x00\x12\x14\n\noneof_bool\x18\x41 \x01(\x08H\x00\x12\x16\n\x0coneof_double\x18\x42 \x01(\x01H\x00\x12\x15\n\x0boneof_float\x18\x43 \x01(\x02H\x00\x12\x16\n\x0coneof_string\x18\x44 \x01(\tH\x00\x12\x15\n\x0boneof_bytes\x18\x45 \x01(\x0cH\x00\x12%\n\noneof_enum\x18\x46 \x01(\x0e\x32\x0f.A.B.C.TestEnumH\x00\x12\'\n\toneof_msg\x18G \x01(\x0b\x32\x12.A.B.C.TestMessageH\x00\x12\x38\n\x0enested_message\x18P \x01(\x0b\x32 .A.B.C.TestMessage.NestedMessage\x1a\x1c\n\rNestedMessage\x12\x0b\n\x03\x66oo\x18\x01 \x01(\x05\x42\n\n\x08my_oneof*,\n\x08TestEnum\x12\x0b\n\x07\x44\x65\x66\x61ult\x10\x00\x12\x05\n\x01\x41\x10\x01\x12\x05\n\x01\x42\x10\x02\x12\x05\n\x01\x43\x10\x03"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError => e
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
    ["A.B.C.TestImportedMessage", "ruby_generated_code_proto2_import.proto"],
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module A
  module B
    module C
      TestMessage = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("A.B.C.TestMessage").msgclass
      TestMessage::NestedMessage = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("A.B.C.TestMessage.NestedMessage").msgclass
      TestEnum = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("A.B.C.TestEnum").enummodule
    end
  end
end
